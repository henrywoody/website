<template>
    <main>
        <div class="welcome">
            <h1 @click="this.reset">welcome</h1>
        </div>
        <canvas id="welcome-canvas"></canvas>
    </main>
</template>

<script>
export default {
    name: "Index",
    mounted() {
        this.setUp();
        this.run();
    },
    beforeDestroy() {
        window.cancelAnimationFrame(this.requestId);
    },
    methods: {
        setUp() {
            this.canvas = document.querySelector("#welcome-canvas");
            this.canvas.width = window.innerWidth;
            this.canvas.height = window.innerHeight;
            this.width = this.canvas.width;
            this.height = this.canvas.height;
            this.ctx = this.canvas.getContext("2d");
            this.requestId = null;

            this.stepSize = 1;
            this.splitProb = this.stepSize/2 * 0.015;
            this.numPointChoices = [2,4,6,10,20];
            this.angleChoices = [2,3,4,6,10].map(e => Math.PI/e);
            this.splitAngleDiff = this.angleChoices[Math.round(Math.random() * (this.angleChoices.length - 1))];
            this.pointSets = [];
            this.maxPointSetSize = Math.pow(2, 7);
            this.maxAge = 5000;

            this.counter = 0;
            this.phaseDuration = 750;
            const isDarkModeOn = window.localStorage.getItem("isDarkModeOn") === "true";
            if (isDarkModeOn) {
                this.phases = ["#ffffff", "#ffd700"];
                this.fadeFactor = 1.7;
            } else {
                this.phases = ["#222222", "#16b2ff"];
                this.fadeFactor = 1.5;
            }
            this.phaseIndex = 0;
            this.maxPhases = 2;
            this.totalPhases = 0;
            this.refreshTime = Math.round(Math.log(this.stepSize)) + 1;

            this.addPointSet();
        },

        reset() {
            window.cancelAnimationFrame(this.requestId);
            this.setUp();
            this.run();
        },

        run() {
            if (this.counter % this.refreshTime == 0) {
                this.maybeSplitPoints();
                this.movePoints();
                this.drawPoints();
            }

            if (this.totalPhases < this.maxPhases && this.counter++ >= this.phaseDuration) {
                this.addPointSet();
                this.counter = 0;
            }

            this.removeOldPointSets();
            this.requestId = window.requestAnimationFrame(this.run);
        },

        addPointSet() {
            this.pointSets.push({
                id: this.phaseIndex,
                alpha: 1,
                age: 0,
                points: this.setPoints()
            });
            this.totalPhases++;
            this.phaseIndex = (this.phaseIndex + 1) % this.phases.length;
        },

        setPoints() {
            const numPoints = this.numPointChoices[Math.floor(Math.random() * this.numPointChoices.length)];
            const points = [];
            for (let i = 0; i < numPoints; i++) {
                points.push({
                    x: this.width/2,
                    y: this.height/2,
                    d: 2*i * Math.PI/numPoints
                });
            }
            return points;
        },

        movePoints() {
            this.pointSets = this.pointSets.map(pointSet => {
                return {
                    ...pointSet,
                    age: ++pointSet.age,
                    points: pointSet.points.map(point => ({
                        x: point.x + this.stepSize * Math.cos(point.d),
                        y: point.y + this.stepSize * Math.sin(point.d),
                        d: point.d
                    }))
                }
            });
        },

        maybeSplitPoints() {
            this.pointSets = this.pointSets.map(pointSet => {
                if (pointSet.points.length <= this.maxPointSetSize && Math.random() < this.splitProb) {
                    return {
                        ...pointSet,
                        alpha: pointSet.alpha / this.fadeFactor,
                        points: pointSet.points.reduce((a,e) => a.concat([
                            {x: e.x, y: e.y, d: e.d + this.splitAngleDiff},
                            {x: e.x, y: e.y, d: e.d - this.splitAngleDiff}
                        ]), [])
                    }
                } else {
                    return pointSet;
                }
            });
        },

        drawPoints() {
            for (const pointSet of this.pointSets) {
                this.ctx.fillStyle = this.phases[pointSet.id];
                this.ctx.globalAlpha = pointSet.alpha;
                for (const point of pointSet.points) {
                    this.ctx.fillRect(point.x, point.y, this.stepSize, this.stepSize);
                }
            }
        },

        removeOldPointSets() {
            this.pointSets = this.pointSets.filter(e => e.age <= this.maxAge);
        },
    }
}
</script>

<style scoped>
main {
    max-width: 100vw;
    height: calc(100vh - var(--header-height));
}

.welcome {
    position: absolute;
    top: 40%;
    left: 10%;
    margin: 0 auto;
    text-align: center;
}

.welcome h1 {
    font-family: Roboto, Arial;
    font-weight: 300;
    font-size: 3.5rem;
    letter-spacing: .3rem;
    color: var(--color);
    cursor: pointer;
    text-shadow:
        -1px -1px 0 var(--background-color),
        1px -1px 0 var(--background-color),
        -1px 1px 0 var(--background-color),
        1px 1px 0 var(--background-color);

    transition: var(--dark-mode-transition), var(--dark-mode-transition-time) text-shadow;
}

canvas {
    width: 100vw;
    height: 100vh;
    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;
}
</style>