<template>
    <CanvasProject title="Two Dimensional Cellular Automata" :setUp="setUp" :update="update">
        <template slot="controls">
            <form class="controls" @submit.prevent>
                <div class="input-container">
                    <label for="speed">Speed</label>
                    <input id="speed" type="range" name="speed" max="10" min="1" v-model="speed">
                    <input class="range-standin" type="number" name="speed" max="10" min="1" v-model="speed">
                </div>

                <div class="input-container">
                    <label for="alive-color">Alive Color</label>
                    <input type="color" name="alive-color" v-model="aliveColor">
                </div>

                <div class="input-container">
                    <label for="dead-color">Dead Color</label>
                    <input type="color" name="dead-color" v-model="deadColor">
                </div>

                <div class="input-container">
                    <button type="button" @click="reset">Reset</button>
                </div>
            </form>
        </template>

        <template slot="description">
            <p>
                A two dimensional cellular automaton is a discrete mathematical model for simulating complex systems. The environment is a 2D grid of cells where each cell can have one several possible states.
            </p>

            <p>
                The state of each cell in the next stage, or time step, is determined by its current state and the states of its neighbors. How neighbors are selected can vary, but the two most common choices are the Von Neumann neighborhood (4 adjacent neighbors) and the Moore neighborhood (8 surrounding cells).
            </p>

            <p>
                The next state of a cell is determined according to a rule. A rule is a mapping from arrangements of current cell states to the next cell state. Generally, rules in the two dimensional case consider only the number of neighboring cells with each state, rather than the actual spacial arrangement of those cells and states. For example, a rule can specify that any neighbor with three living neighbors that is currently alive stays alive, and dies otherwise.
            </p>

            <h2>Game of Life</h2>

            <p>
                <a href="https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life">Conway's Game of Life</a> is a classic rule for two dimensional cellular automata. There are two possible states, generally interpreted as live and dead, and the system uses the Moore neighborhood. The rules are as follows:

                <ol>
                    <li>Any living cell with fewer than two live neighbors dies from isolation</li>
                    <li>Any living cell with two or three live neighbors stays alive</li>
                    <li>Any dead cell with exactly three live neighbors comes to life</li>
                    <li>Any living cell with more than three live neighbors dies from overcrowding</li>
                </ol>
            </p>
        </template>
    </CanvasProject>
</template>

<script>
import CanvasProject from './CanvasProject.vue';

export default {
    name: "TDCA",
    components: {
        CanvasProject,
    },
    data() {
        return {
            aliveColor: "#000000",
            deadColor: null,
            speed: 3,
        }
    },
    methods: {
        setUp(canvas) {
            this.canvas = canvas;
            this.ctx = this.canvas.getContext("2d");
            this.cellSize = 5;
            this.initalLiveProb = 0.15;
            this.grid = this.getRandomGrid(this.initalLiveProb);
            this.tickCounter = 0;
        },

        update() {
            this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
            this.updateGrid();
            this.draw();
        },

        reset() {
            this.setUp(this.canvas);
        },

        updateGrid() {
            this.tickCounter = (this.tickCounter + 1) % (11 - this.speed);
            if (this.tickCounter === 0) {
                const newGrid = this.getBlankGrid();
                for (let i = 0; i < this.grid.length; i++) {
                    for (let j = 0; j < this.grid[0].length; j++) {
                        newGrid[i][j] = this.getNextCellState(i, j);
                    }
                }
                this.grid = newGrid;
            }
        },

        getNextCellState(i, j) {
            const cellState = this.grid[i][j];
            let neighborhood = this.grid.slice(i - 1, i + 2).map(e => e.slice(j - 1, j + 2));
            neighborhood = neighborhood.reduce((a, e) => a.concat(e), []);
            const numLiving = neighborhood.reduce((a, e) => a + e, 0);

            switch (numLiving - cellState) {
            case 2:
                return cellState;
            case 3:
                return 1;
            default:
                return 0;
            }
        },

        getRandomGrid(liveProb) {
            liveProb = liveProb || 0.5;
            const grid = this.getBlankGrid();
            for (let i = 0; i < grid.length; i++) {
                for (let j = 0; j < grid[0].length; j++) {
                    grid[i][j] = Math.random() < liveProb ? 1 : 0;
                }
            }
            return grid;
        },

        getBlankGrid() {
            const nRows = this.getNumRows();
            const nCols = this.getNumCols();
            const grid = Array(nRows).fill(0);
            for (let i = 0; i < nRows; i++) {
                grid[i] = Array(nCols).fill(0);
            }
            return grid;
        },

        getNumRows() {
            return Math.floor(this.canvas.height / this.cellSize);
        },

        getNumCols() {
            return Math.floor(this.canvas.width / this.cellSize);
        },

        draw() {
            for (let i = 0; i < this.grid.length; i++) {
                for (let j = 1; j < this.grid[0].length; j++) {
                    const state = this.grid[i][j];
                    const color = this.getColor(state);
                    if (color) {
                        const y = i * this.cellSize;
                        const x = j * this.cellSize;
                        this.ctx.fillStyle = color;
                        this.ctx.fillRect(x, y, this.cellSize, this.cellSize);
                    }
                }
            }
        },

        getColor(state) {
            switch (state) {
            case 1:
                return this.aliveColor;
            case 0:
                return this.deadColor;
            }
        },
    },
}
</script>

<style scoped>
button {
    margin-top: .5rem;
}
</style>