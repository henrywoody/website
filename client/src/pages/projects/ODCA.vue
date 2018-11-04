<template>
    <CanvasProject title="One Dimensional Cellular Automata" :setUp="setUp" :update="update">
        <template slot="controls">
            <form class="controls" @submit.prevent>
                <div class="input-container">
                    <label for="rule">Rule</label>
                    <input id="rule" type="number" name="rule" max="255" min="0" v-model="rule">
                </div>

                <div class="input-container">
                    <label for="random-rules">Random Rules</label>
                    <input id="random-rules" type="checkbox" name="random-rules" v-model="randomRules">
                </div>

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
                A one dimensional cellular automaton is a discrete mathematical model for simulating complex systems. The environment is a linear strip of cells where each cell can have one of several possible states. The simplest interesting example, is a binary state cellular automaton, for example alive or dead. The state of each cell in the next stage, or time step, is determined by its current state and the states of the cells on either side of it.
            </p>

            <p>
                How the configuration of the three cells (self and neighbors) impact a given cell depends on the selected rule. A rule is a mapping from "parent" cell configurations to cell states. For example, a rule might specify that a cell is alive in the next stage if it and its neighbors are all currently alive and is dead otherwise.
            </p>

            <h2>Elementary Cellular Automata</h2>

            <p>
                An elementary cellular automaton is a binary state (e.g. 0 and 1), one dimenstional cellular automata. With two possible states for each cell and three "parent" cells, there are 2<sup>3</sup> possible parent configurations. Here they are:
            </p>

            <p class="parent-configs">
                <span>111</span><span>110</span><span>101</span><span>100</span><span>011</span><span>010</span><span>001</span><span>000</span>
            </p>

            <p>
                Each of these configurations can lead to a live cell or dead cell. Thus there are 2<sup>2<sup>3</sup></sup> = 256 rules for elementary cellular automata. Each rule is basically just a number that, when written in 8-bit binary form, determines how the parent configurations map to states for each cells. For example, rule 30 (or 00011110), looks like this:
            </p>

            <p class="parent-configs">
                 <span>111</span><span>110</span><span>101</span><span>100</span><span>011</span><span>010</span><span>001</span><span>000</span>
            </p>
            <p class="child-configs">
                <span>0</span><span>0</span><span>0</span><span>1</span><span>1</span><span>1</span><span>1</span><span>0</span>
            </p>

            <p>
                Where the top row is the parent configuration and the bottom row is the state of the cell with that parent configuration in the next stage.
            </p>

            <p>
                See this <a href="http://mathworld.wolfram.com/ElementaryCellularAutomaton.html">WolframMathWorld article</a> for more.
            </p>
        </template>
    </CanvasProject>
</template>

<script>
import CanvasProject from './CanvasProject.vue';

export default {
    name: "ODCA",
    components: {
        CanvasProject
    },
    data() {
        return {
            rule: null,
            aliveColor: "#000000",
            deadColor: null,
            randomRules: true,
            speed: 6,
        }
    }, 
    methods: {
        setUp(canvas) {
            this.canvas = canvas;
            this.ctx = this.canvas.getContext("2d");
            this.cellSize = 5;
            this.ruleChangeFrequency = 200;
            this.grid = [];
            this.stageCounter = 0;
            this.tickCounter = 0;

            const goodStartingRules = ["30", "57", "150", "182"];
            this.rule = goodStartingRules[Math.floor(Math.random() * goodStartingRules.length)];

            this.setFirstRow();
        },

        update() {
            const speed = Math.min(Math.max(Number(this.speed), 1), 10);
            this.tickCounter = (this.tickCounter + 1) % (11 - speed);
            if (this.tickCounter === 0) {
                this.stageCounter = (this.stageCounter + 1) % this.ruleChangeFrequency;
                this.ctx.clearRect(0,0,this.canvas.width, this.canvas.height);
                this.generateRow();
                this.maybeTrimRows();
                this.draw();
                this.maybeChangeRule();
            }
        },

        reset() {
            this.setUp(this.canvas);
        },

        setFirstRow() {
            const numCells = this.getNumCellsPerRow();
            const row = Array(numCells).fill(0);
            row[Math.floor(numCells / 2)] = 1;
            this.grid.push(row);
        },

        generateRow() {
            const numCells = this.getNumCellsPerRow();
            const prevRow = this.grid.slice(-1)[0];
            const nextRow = Array(numCells).fill(0);
            const ruleMap = this.getRuleMap();

            for (let i = 0; i < prevRow.length; i++) {
                const parentConfig = this.getParentConfig(i, prevRow);
                nextRow[i] = ruleMap[parentConfig];
            }
            this.grid.push(nextRow);
        },

        getRuleMap() {
            const parentConfigs = Array(8).fill(0).map((_, i) => toBinary(i, 3)).reverse();
            const rule = Math.min(Math.max(Number(this.rule), 0), 255);
            const outcomes = toBinary(rule, 8, true);
            const ruleMap = {};

            for (let i = 0; i < parentConfigs.length; i++) {
                ruleMap[parentConfigs[i]] = outcomes[i];
            }
            return ruleMap;
        },

        getNumCellsPerRow() {
            return Math.floor(this.canvas.width / this.cellSize);
        },

        getParentConfig(i, prevRow) {
            let parentConfig;
            switch (i) {
            case 0:
                parentConfig = [0].concat(prevRow.slice(0,2));
                break;
            case prevRow.length - 1:
                parentConfig = prevRow.slice(-2).concat([0]);
                break;
            default:
                parentConfig = prevRow.slice(i-1, i+2);
            }
            return parentConfig.join("");
        },

        maybeTrimRows() {
            const numRows = Math.floor(this.canvas.height / this.cellSize);
            if (this.grid.length > numRows) {
                const numToTrim = this.grid.length - numRows;
                this.grid = this.grid.slice(numToTrim);
            }
        },

        maybeChangeRule() {
            if (this.randomRules && this.stageCounter === 0) {
                this.rule = Math.floor(Math.random() * 256).toString();
            }
        },

        draw() {
            for (let i = 0; i < this.grid.length; i++) {
                for (let j = 0; j < this.grid.slice(-1)[0].length; j++) {
                    const y = i * this.cellSize;
                    const x = j * this.cellSize;
                    const cellState = this.grid[i][j];
                    const color = cellState === 1 ? this.aliveColor : this.deadColor;
                    if (!color) {
                        continue;
                    }
                    this.ctx.fillStyle = color
                    this.ctx.fillRect(x, y, this.cellSize, this.cellSize);
                }
            }
        },
    }
}

function toBinary(num, length, asArray) {
    const binary = num.toString(2);
    const padding = "0".repeat(length - binary.length);
    const fullBinary = padding + binary;
    if (asArray) {
        return fullBinary.split("").map(e => Number(e));
    }
    return fullBinary;
}
</script>

<style scoped>
button {
    margin-top: .5rem;
}

.parent-configs, .child-configs {
    display: flex;
    flex-flow: row;
    justify-content: space-between;
}

.child-configs {
    padding: 0 .5rem;
}
</style>