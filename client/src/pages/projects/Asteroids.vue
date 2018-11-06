<template>
    <div>
        <canvas></canvas>
        <div v-if="gameOver" class="game-over">
            <span>Game Over</span>
            <button @click="reset">Play Again?</button>
        </div>
    </div>
</template>

<script>
export default {
    name: "Asteroids",

    mounted() {
        this.setUpCanvas();
        this.reset();
        this.run();
        window.addEventListener("resize", this.resizeCanvas);
        window.addEventListener("keyup", this.handleKeyUp);
        window.addEventListener("keydown", this.handleKeyDown);
    },
    beforeDestroy() {
        window.cancelAnimationFrame(this.requestId);
        window.removeEventListener("resize", this.resizeCanvas);
        window.removeEventListener("keyup", this.handleKeyUp);
        window.removeEventListener("keydown", this.handleKeyDown);
    },
    data() {
        return {
            gameOver: false
        }
    },

    methods: {
        setUpCanvas() {
            this.resizeCanvas();
            this.ctx = this.canvas.getContext("2d");
            this.requestId = null;
        },
        reset() {
            const middleOfCanvas = [this.canvas.width/2, this.canvas.height/2];
            this.ship = new Ship(middleOfCanvas, this.addBullet);
            this.bullets = [];
            this.asteroids = [];
            this.gameOver = false;

            this.time = 0;
        },
        run() {
            this.update();
            this.draw();
            this.requestId = window.requestAnimationFrame(this.run);
        },
        update() {
            const color = document.body.classList.contains("dark") ? "#ffffff" : "#333333";
            this.ctx.fillStyle = color;
            this.ctx.strokeStyle = color;

            this.ship.update(this.canvas.width, this.canvas.height);
            for (const bullet of this.bullets) {
                bullet.update();
            }
            for (const asteroid of this.asteroids) {
                asteroid.update();
            }
            this.removeOutOfBoundsObjects();
            this.makeRandomAsteroids();
            this.checkCollisons();
            this.removeInactiveObjects();

            if (!this.gameOver) {
                this.time++;
            }
        },
        handleKeyUp(event) {
            this.ship.handleUserInput(event, "up")
        },
        handleKeyDown(event) {
            this.ship.handleUserInput(event, "down");
        },
        addBullet(bullet) {
            this.bullets.push(bullet)
        },
        makeRandomAsteroids() {
            const asteroidCreationProb = Math.log(Math.log(this.time)) * 0.01;
            if (Math.random() < asteroidCreationProb) {
                const numToCreate = Math.floor(Math.random() * 3) + 1;
                for (let i = 0; i < numToCreate; i++) {
                    let position, direction;
                    const rand = Math.random();
                    if (rand < 0.25) {
                        position = [this.xMin, Math.random() * this.yMax];
                        direction = Math.random() * Math.PI + Math.PI;
                    } else if (rand < 0.5) {
                        position = [this.xMax, Math.random() * this.yMax];
                        direction = -1 * Math.random() * Math.PI - Math.PI;
                    } else if (rand < 0.75) {
                        position = [Math.random() * this.xMax, this.yMin];
                        direction = Math.random() * Math.PI;
                    } else {
                        position = [Math.random() * this.xMax, this.yMax];
                        direction = -1 * Math.random() * Math.PI;
                    }
                    this.asteroids.push(new Asteroid(position, direction))
                }
            }
        },
        removeOutOfBoundsObjects() {
            const filterFunc = e => {
                const inXBounds = this.xMin <= e.position[0] && e.position[0] <= this.xMax;
                const inYBounds = this.yMin <= e.position[1] && e.position[1] <= this.yMax;
                return inXBounds && inYBounds;
            }
            this.bullets = this.bullets.filter(filterFunc);
            this.asteroids = this.asteroids.filter(filterFunc);
        },
        checkCollisons() {
            this.checkShipAsteroidCollisons();
            this.checkBulletAsteroidCollisons();
        },
        checkShipAsteroidCollisons() {
            for (const asteroid of this.asteroids) {
                const distance = getDistanceBetween(asteroid.position, this.ship.position);
                if (distance <= asteroid.size + this.ship.size/2) {
                    this.ship.crash();
                    this.endGame();
                }
            }
        },
        checkBulletAsteroidCollisons() {
            const newAsteroids = [];
            for (const bullet of this.bullets) {
                for (const asteroid of this.asteroids) {
                    const distance = getDistanceBetween(asteroid.position, bullet.position);
                    if (distance <= asteroid.size + bullet.size) {
                        bullet.hit();
                        newAsteroids.push(...asteroid.split());
                    }
                }
            }
            this.asteroids.push(...newAsteroids);
        },
        removeInactiveObjects() {
            const filterFunc = e => e.isActive;
            this.bullets = this.bullets.filter(filterFunc);
            this.asteroids = this.asteroids.filter(filterFunc);
        },
        endGame() {
            this.gameOver = true;
        },
        draw() {
            this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
            if (!this.gameOver) {
                this.ship.draw(this.ctx);
            }
            for (const bullet of this.bullets) {
                bullet.draw(this.ctx);
            }
            for (const asteroid of this.asteroids) {
                asteroid.draw(this.ctx);
            }
        },
        resizeCanvas() {
            this.canvas = this.$el.querySelector("canvas");
            this.canvas.width = window.innerWidth;
            const headerHeight = document.querySelector(".header-bar").offsetHeight;
            this.canvas.height = window.innerHeight - headerHeight;

            const bounaryExtension = 100;
            this.xMin = -bounaryExtension;
            this.xMax = this.canvas.width + bounaryExtension;
            this.yMin = -bounaryExtension;
            this.yMax = this.canvas.height + bounaryExtension;
        },
    }
}

class Ship {
    constructor(initPositon, addBullet) {
        this.position = initPositon;
        this.direction = 3 * Math.PI/2;
        this.velocity = [0, 0];
        this.addBullet = addBullet;

        this.maxSpeed = 4;
        this.acceleration = 0.15;
        this.deceleration = 0.995;
        this.turnSpeed = Math.PI/30;
        this.size = 12;
        this.shotCoolDownTime = 15;

        this.isAccelerating = false;
        this.isTurningLeft = false;
        this.isTurningRight = false;
        this.shotCounter = 0;

        this.keyActions = {
            37: "LEFT",
            38: "ACCELERATE",
            39: "RIGHT",
            65: "LEFT",
            87: "ACCELERATE",
            68: "RIGHT",
            32: "SHOOT",
        }

        this.isActive = true;
    }

    update(xMax, yMax) {
        if (!this.isActive) {
            return;
        }

        this.move(xMax, yMax);
        if (this.isAccelerating) {
            this.accelerate();
        }
        if (this.isTurningLeft) {
            this.turnLeft();
        }
        if (this.isTurningRight) {
            this.turnRight();
        }
        this.shotCounter = Math.max(this.shotCounter - 1, 0);
        const canShoot = this.shotCounter === 0;
        if (this.isShooting && canShoot) {
            this.shoot();
        }
    }

    move(xMax, yMax) {
        this.position[0] = (this.position[0] + this.velocity[0] + xMax) % xMax;
        this.position[1] = (this.position[1] + this.velocity[1] + yMax) % yMax;
        this.velocity[0] *= this.deceleration;
        this.velocity[1] *= this.deceleration;
    }

    accelerate() {
        const speed = Math.sqrt(Math.pow(this.velocity[0], 2) + Math.pow(this.velocity[1], 2));
        const canAccelerate = speed < this.maxSpeed;
        if (canAccelerate) {
            this.velocity[0] = this.velocity[0] + this.acceleration * Math.cos(this.direction);
            this.velocity[1] = this.velocity[1] + this.acceleration * Math.sin(this.direction);
        }
    }

    turnLeft() {
        this.direction = (this.direction - this.turnSpeed + 2 * Math.PI) % (2 * Math.PI);
    }

    turnRight() {
        this.direction = (this.direction + this.turnSpeed) % (2 * Math.PI);
    }

    shoot() {
        const bullet = new Bullet([...this.position], this.direction);
        this.addBullet(bullet);
        this.shotCounter = this.shotCoolDownTime;
    }

    crash() {
        this.isActive = false;
    }

    handleUserInput(event, type) {
        const action = this.keyActions[event.keyCode];
        const setTo = type === "down";
        switch (action) {
        case "ACCELERATE":
            this.isAccelerating = setTo;
            break;
        case "LEFT":
            this.isTurningLeft = setTo;
            break;
        case "RIGHT":
            this.isTurningRight = setTo;
            break;
        case "SHOOT":
            this.isShooting = setTo;
            break;
        }
    }

    draw(ctx) {
        const frontPoint = [
            this.position[0] + this.size * Math.cos(this.direction),
            this.position[1] + this.size * Math.sin(this.direction)
        ]
        const backSidePoint1 = [
            this.position[0] + this.size/2 * Math.cos(this.direction + 2 * Math.PI/3),
            this.position[1] + this.size/2 * Math.sin(this.direction + 2 * Math.PI/3),
        ]
        const backSidePoint2 = [
            this.position[0] + this.size/2 * Math.cos(this.direction - 2 * Math.PI/3),
            this.position[1] + this.size/2 * Math.sin(this.direction - 2 * Math.PI/3),
        ]
        ctx.beginPath();
        ctx.moveTo(...frontPoint);
        ctx.lineTo(...backSidePoint1);
        ctx.lineTo(...this.position);
        ctx.lineTo(...backSidePoint2);
        ctx.lineTo(...frontPoint);
        ctx.stroke();

    }
}

class Bullet {
    constructor(initPositon, direction) {
        this.position = initPositon;
        this.speed = 10;
        this.velocity = [
            this.speed * Math.cos(direction),
            this.speed * Math.sin(direction)
        ];
        this.size = 1;

        this.isActive = true;
    }

    update() {
        this.move();
    }

    move() {
        this.position[0] += this.velocity[0];
        this.position[1] += this.velocity[1];
    }

    hit() {
        this.isActive = false;
    }

    draw(ctx) {
        ctx.beginPath();
        ctx.arc(...this.position, this.size, 0, 2 * Math.PI);
        ctx.stroke();
    }
}

class Asteroid {
    constructor(initPositon, direction, tier) {
        this.position = initPositon;

        const maxAsteroidSpeed = 4;
        const minAsteroidSpeed = 1;
        this.speed = Math.random() * (maxAsteroidSpeed - minAsteroidSpeed) + minAsteroidSpeed;
        this.velocity = [
            this.speed * Math.cos(direction),
            this.speed * Math.sin(direction)
        ];

        const numTiers = 5;
        this.tier = tier || Math.floor(Math.random() * numTiers);
        this.size = 10 * Math.pow(this.tier, 1.5);

        this.isActive = true;
    }

    update() {
        this.move();
    }

    move() {
        this.position[0] += this.velocity[0];
        this.position[1] += this.velocity[1];
    }

    split() {
        this.isActive = false;
        const numToCreate = 3;
        if (this.tier > 1) {
            return Array(numToCreate).fill(0).map(_ => {
                return new Asteroid(
                    [...this.position],
                    Math.random() * 2 * Math.PI,
                    this.tier - 1
                )
            });
        } else {
            return [];
        }
    }

    draw(ctx) {
        ctx.beginPath();
        ctx.arc(...this.position, this.size, 0, 2 * Math.PI);
        ctx.stroke();
    }
}

function getDistanceBetween(a, b) {
    const dx = b[0] - a[0]
    const dy = b[1] - a[1];
    return Math.sqrt(Math.pow(dx, 2) + Math.pow(dy, 2));
}
</script>

<style scoped>
canvas {
    position: absolute;
    left: 0;
    top: var(--header-height);
}

.game-over {
    --width: 7rem;
    --height: 4rem;
    width: var(--width);
    height: var(--height);
    display: flex;
    flex-flow: column;
    justify-content: center;
    position: absolute;
    left: calc((100vw - var(--width)) / 2);
    top: calc((100vh - var(--height)) / 2);
    z-index: 100;
}

.game-over span {
    font-size: 1.25rem;
    text-align: center;
    margin-bottom: .5rem;
    display: block;
}

</style>