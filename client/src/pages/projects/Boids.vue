<template>
    <CanvasProject title="Boids" :setUp="setUp" :update="update">
        <template slot="controls">
            <form class="controls" @submit.prevent>
                <div class="input-container">
                    <label for="numBoids">Population</label>
                    <input id="numBoids" type="number" name="numBoids" v-model="numBoids" @change="updateBoidPopulation">
                </div>

                <div class="input-container">
                    <label for="flocking">Flocking</label>
                    <input type="checkbox" name="flocking" v-model="isFlockingOn">
                </div>

                <div class="input-container">
                    <label for="colors">Colors</label>
                    <input type="checkbox" name="colors" v-model="areColorsOn">
                </div>
            </form>
        </template>
        <template slot="description">
            <p>
                Boids (bird-oids) are artificial creatures that simulate the flocking behavior of natural birds. They were originally described by Craig W. Reynolds in a <a href="http://www.cs.toronto.edu/~dt/siggraph97-course/cwr87/">1987 paper</a> on simulating aggregate behavior of groups of individuals.
            </p>

            <p>
                Each boid independently decides how to move according to the following set of rules:

                <ol>
                    <li>Move toward the center of mass of the other boids within some neighborhood</li>
                    <li>Keep a small distance from other boids to avoid collisions</li>
                    <li>Try to match the velocity of nearby boids</li>
                </ol>
            </p>

            <p>
                I've added the 'colors' flocking rule here, which incorporates chromatic distance, in addition to spacial, when each boids considers their neighbors. So in order to be considered a neighbor, a boid must be relatively close spacially and have a similar color. Basically boids will only flock with similarly colored boids if 'colors' is on.
            </p>
        </template>
    </CanvasProject>
</template>

<script>
import CanvasProject from './CanvasProject.vue';

export default {
    name: "Boids",
    components: {
        CanvasProject
    },
    data() {
        return {
            isFlockingOn: true,
            areColorsOn: false,
            numBoids: 100,
        }
    },
    methods: {
        setUp(canvas) {
            this.canvas = canvas;
            this.ctx = this.canvas.getContext("2d");
            this.maxSpeed = 2;
            this.boids = [];
            this.addBoids(100);
        },

        update() {
            this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
            for (const boid of this.boids) {
                boid.updatePosition(this.boids, this.canvas, this.isFlockingOn, this.areColorsOn);
                boid.draw(this.ctx);
            }
        },

        updateBoidPopulation() {
            if (this.numBoids > this.boids.length) {
                this.addBoids(this.numBoids - this.boids.length);
            } else if (this.boids.length > this.numBoids) {
                this.boids = this.boids.slice(this.boids.length - this.numBoids);
            }
        },

        addBoids(num) {
            for (let i = 0; i < num; i++) {
                const color = this.getRandomColor();
                const initPosition = [
                    Math.random() * this.canvas.width,
                    Math.random() * this.canvas.height
                ];
                const initVelocity = [
                    Math.random() * (this.maxSpeed - 1) + 1,
                    Math.random() * 2 * Math.PI
                ];
                this.boids.push(
                    new Boid(color, initPosition, initVelocity)
                );
            }
        },

        getRandomColor(rRange=[0,255], gRange=[0,255], bRange=[0,255]) {
            return "#" + [rRange, gRange, bRange].map(e => {
                return (Math.round(Math.random() * (e[1] - e[0])) + e[0]).toString(16);
            }).map(e => {
                return e.length < 2 ? '0' + e : e;
            }).join('');
        }
    }
}

class Boid {
    constructor(color, initPosition, initVelocity) {
        this.color = color;
        this.position = initPosition;
        this.velocity = initVelocity || [0,0]; // [r, theta]
        this.size = 3;
        this.beakLength = 1.5 * this.size;

        this.neighborhoodDistance = 30;
        this.colorTolerance = 90;
        this.avoidDistance = this.size * 3;
        this.thisSpeedInfluence = 0;
        this.thisAngleInfluence = 1;
        this.centerOfMassInfluence = 0.75;
        this.avoidSpeedInfluence = 1;
        this.avoidAngleInfluence = 1;
        this.neighborSpeedInfluence = 100;
        this.neighborAngleInfluence = 100;
    }

    updatePosition(otherBoids, canvas, isFlockingOn, areColorsOn) {
        this.updateVelocity(otherBoids, isFlockingOn, areColorsOn);

        this.position[0] += this.velocity[0] * Math.cos(this.velocity[1]);
        this.position[1] += this.velocity[0] * Math.sin(this.velocity[1]);

        // Torus
        this.position[0] = (this.position[0] + canvas.width) % canvas.width;
        this.position[1] = (this.position[1] + canvas.height) % canvas.height;
    }

    updateVelocity(otherBoids, isFlockingOn, areColorsOn) {
        const closeBoids = this.getNeighborsAndDistances(otherBoids);
        const boidsToAvoid = closeBoids.filter(e => e[1] <= this.avoidDistance);
        this.checkAndHandleCollisions(boidsToAvoid);

        const neighborDistances = closeBoids.filter(e => {
            const colorDistance = getColorDistanceBetween(this.color, e[0].color);
            return !areColorsOn || colorDistance < this.colorTolerance;
        });
        if (!neighborDistances.length) return;

        // Rules
        const angleToCenter = isFlockingOn ? this.getAngleToCenterOfMass(neighborDistances) : this.velocity[1];
        const avoidVelocity = this.getAvoidVelocity(boidsToAvoid);
        const aveVelocity = isFlockingOn ? this.getAveVelocity(neighborDistances) : this.velocity;

        let speedSum = this.thisSpeedInfluence * this.velocity[0] + this.neighborSpeedInfluence * aveVelocity[0];
        let speedWeightSum = this.thisSpeedInfluence + this.neighborSpeedInfluence;

        const angles = [this.velocity[1], aveVelocity[1], angleToCenter];
        const angleWeights = [this.thisAngleInfluence, this.neighborAngleInfluence, this.centerOfMassInfluence];
        if (avoidVelocity) {
            speedSum += this.avoidSpeedInfluence * aveVelocity[0];
            speedWeightSum += this.avoidSpeedInfluence;

            angles.push(aveVelocity[1] + Math.PI); // opposite direction
            angleWeights.push(this.avoidAngleInfluence);
        }

        const randAngle = Math.random() * 0.1 - 0.05;

        const newSpeed = speedSum / speedWeightSum;
        const newAngle = ((getAveAngle(angles, angleWeights) + randAngle) + 2 * Math.PI) % (2 * Math.PI);

        this.velocity = [newSpeed, newAngle];
    }

    getNeighborsAndDistances(otherBoids) {
        return otherBoids.reduce((a,e) => {
            const distance = getDistanceBetween(this.position, e.position);
            const isNeighbor = distance <= this.neighborhoodDistance && e !== this;
            return isNeighbor ? a.concat([[e, distance]]) : a;
        }, []);
    }

    // Rule 1
    getAngleToCenterOfMass(neighborDistances) {
        const positionSum = neighborDistances.reduce((a, e) => {
            return [a[0] + e[0].position[0], a[1] + e[0].position[1]];
        }, [0,0]);

        const center = positionSum.map(e => e / neighborDistances.length);

        return getAngleBetween(this.position, center);
    }

    // Rule 2
    getAvoidVelocity(boidsToAvoid) {
        if (!boidsToAvoid.length) return;
        return this.getAveVelocity(boidsToAvoid);
    }

    // Rule 3
    getAveVelocity(neighborDistances) {
        let speedSum = 0;
        let neighborAngles = [];
        for (const neighborDist of neighborDistances) {
            const boid = neighborDist[0];
            speedSum += boid.velocity[0];
            neighborAngles.push(boid.velocity[1]);
        }
        const speed = speedSum / neighborDistances.length;
        const angle = getAveAngle(neighborAngles);
        return [speed, angle];
    }

    checkAndHandleCollisions(boidsToAvoid) {
        for (const [otherBoid, distance] of boidsToAvoid) {
            if (distance <= this.size + otherBoid.size) {
                this.bounceBack(otherBoid)
            }
        }
    }

    bounceBack(otherBoid) {
        const bounceDist = 0.2;
        this.position[0] += otherBoid.position[0] - this.position[0] > 0 ? -bounceDist : bounceDist;
        this.position[1] += otherBoid.position[1] - this.position[1] > 0 ? -bounceDist : bounceDist;
        otherBoid.position[0] += this.position[0] - otherBoid.position[0] > 0 ? -bounceDist : bounceDist;
        otherBoid.position[1] += this.position[1] - otherBoid.position[1] > 0 ? -bounceDist : bounceDist;
    }

    draw(ctx) {
        ctx.fillStyle = this.color;
        ctx.strokeStyle = this.color;
        ctx.beginPath();
        ctx.arc(...this.position, this.size, 0, 2 * Math.PI);
        ctx.fill();
        ctx.closePath();
        ctx.beginPath();
        ctx.moveTo(...this.position);
        ctx.lineTo(
            this.beakLength * Math.cos(this.velocity[1]) + this.position[0],
            this.beakLength * Math.sin(this.velocity[1]) + this.position[1]
        );
        ctx.stroke();
    }
}

function getAngleBetween(a, b) {
    const dx = b[0] - a[0];
    const dy = b[1] - a[1];
    return (Math.atan2(dy, dx) + 2 * Math.PI) % (2 * Math.PI);
}

function getDistanceBetween(a, b) {
    const dx = b[0] - a[0];
    const dy = b[1] - a[1];
    return Math.sqrt(Math.pow(dx, 2) + Math.pow(dy, 2));
}

function getColorDistanceBetween(a, b) {
    const aColor = a.slice(1);
    const bColor = b.slice(1);

    let diff = 0;
    for (let i = 0; i < 3; i++) {
        const aComp = parseInt(aColor.slice(2 * i, 2 * (i + 1)), 16);
        const bComp = parseInt(bColor.slice(2 * i, 2 * (i + 1)), 16);
        diff += Math.abs(aComp - bComp);
    }
    return diff;
}

function getAveAngle(angles, weights) {
    if (angles.length === 1) {
        return angles[0];
    }
    weights = weights || Array(angles.length).fill(1);

    let xSum = 0;
    let ySum = 0;
    for (let i = 0; i < angles.length; i++) {
        xSum += Math.cos(angles[i]) * weights[i];
        ySum += Math.sin(angles[i]) * weights[i];
    }

    const weightSum = weights.reduce((a,e) => a + e);
    const xAve = xSum / weightSum;
    const yAve = ySum / weightSum;
    const aveAngle = Math.atan2(yAve, xAve);
    return aveAngle;
}
</script>
