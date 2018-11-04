<template>
    <main>
        <h1>{{ title }}</h1>

        <div class="canvas-container">
            <canvas></canvas>
        </div>

        <slot name="controls"/>

        <section class="description">
            <slot name="description"/>
        </section>
    </main>
</template>

<script>
export default {
    name: "CanvasProject",
    props: {
        title: String,
        setUp: Function,
        run: Function,
    },
    mounted() {
        this.superSetUp();
        this.superRun();

        window.addEventListener("resize", this.resizeCanvas);
    },
    beforeDestroy() {
        window.cancelAnimationFrame(this.requestId);
        window.removeEventListener("resize", this.resizeCanvas);
    },
    methods: {
        superSetUp() {
            this.resizeCanvas();
            this.ctx = this.canvas.getContext("2d");
            this.requestId = null;

            this.setUp(this.canvas, this.ctx);
        },

        resizeCanvas() {
            this.canvas = this.$el.querySelector("canvas");
            this.canvas.width = window.innerWidth;
            this.canvas.height = window.innerHeight;
        },

        superRun() {
            this.run();
            this.requestId = window.requestAnimationFrame(this.superRun);
        }
    }
}
</script>

<style scoped>
.canvas-container {
    height: 100vh;
}
canvas {
    position: absolute;
    left: 0;
}

p {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}

form.controls {
    width: 50%;
    margin: 1rem auto 2rem;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
}

form.controls .input-container {
    display: flex;
    flex-flow: column;
}

form.controls label {
    text-transform: uppercase;
    letter-spacing: .05rem;
    font-size: .7rem;
    margin-bottom: .1rem;
}
</style>