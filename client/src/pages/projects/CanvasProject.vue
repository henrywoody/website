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
        update: Function,
    },
    mounted() {
        this.superSetUp();
        this.run();

        window.addEventListener("resize", this.resizeCanvas);
    },
    beforeDestroy() {
        window.cancelAnimationFrame(this.requestId);
        window.removeEventListener("resize", this.resizeCanvas);
    },
    methods: {
        superSetUp() {
            this.resizeCanvas();
            this.requestId = null;

            this.setUp(this.canvas);
        },

        resizeCanvas() {
            this.canvas = this.$el.querySelector("canvas");
            this.canvas.width = window.innerWidth;
            this.canvas.height = window.innerHeight;
        },

        run() {
            this.update();
            this.requestId = window.requestAnimationFrame(this.run);
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

section.description {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}

form.controls {
    width: 90%;
    max-width: 30rem;
    margin: 1rem auto 1rem;
    display: flex;
    flex-flow: row;
    flex-wrap: wrap;
    justify-content: space-between;
}

form.controls .input-container {
    height: 100%;
    display: flex;
    flex-flow: column;
    margin-bottom: 1rem;
}

form.controls .input-container:not(:first-of-type) {
    margin-left: .5rem;
}

form.controls .input-container:not(:last-of-type) {
    margin-right: .5rem;
}

form.controls label {
    text-transform: uppercase;
    letter-spacing: .05rem;
    font-size: .7rem;
    margin-bottom: .1rem;
}
</style>