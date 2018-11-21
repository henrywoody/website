<template>
    <div class="leaderboard">
        <div class="score-info">
            <span v-if="isNewHighScore" class="new-high-score">New High Score!</span>
            <span class="game-over-score">Score: {{ score }}</span>
        </div>
        <div v-if="highScores.length">
            <span>Leaderboard</span>
            <ol>
                <li v-for="(highScore, i) in highScores" :key="i">
                    <span v-if="highScore.name === '$this'">
                        <span>{{ highScore.score }}</span>
                        <form @submit="postScore">
                            <input ref="playerNameInput"
                                :value="playerName"
                                @input="playerName = $event.target.value.toUpperCase()"
                                :readonly="scoreWasPosted"
                                type="text" size="3" maxlength="3"
                            />
                        </form>
                    </span>
                    <span v-else>
                        <span>{{ highScore.score }}</span>
                        <span>{{ highScore.name }}</span>
                    </span>
                </li>
            </ol>
        </div>    
    </div>
</template>

<script>
export default {
    name: "AsteroidsLeaderboard",

    props: {
        score: Number,
        numScores: Number,
    },

    data() {
        return {
            isNewHighScore: false,
            highScores: [],
            playerName: "",
            scoreWasPosted: false,
        }
    },

    async created() {
        await this.fetchScores();
        this.checkAndHandleNewHighScore();
    },

    beforeDestroy() {
        this.removeFocusForcerFromPlayerNameInput();
    },

    methods: {
        async fetchScores() {
            const response = await fetch(`/api/asteroids_scores?limit=${this.numScores}`);
            this.highScores = await response.json();
        },

        checkAndHandleNewHighScore() {
            if (this.highScores.length && this.score > this.highScores[this.highScores.length - 1].score) {
                this.isNewHighScore = true;
                const insertedHighScores = [...this.highScores];
                let insertIndex;
                if (this.score > this.highScores[0].score) {
                    insertIndex = 0;
                } else {
                    for (let i = this.highScores.length - 1; i >= 0; i--) {
                        if (this.highScores[i].score >= this.score) {
                            insertIndex = i + 1
                            break;
                        }
                    }
                }
                insertedHighScores.splice(insertIndex, 0, {score: this.score, name: "$this"});
                if (insertedHighScores.length > this.numScores) {
                    insertedHighScores.pop();
                }
                this.highScores = insertedHighScores;
            } else if (this.highScores.length < this.numScores && this.score > 0) {
                this.isNewHighScore = true;
                this.highScores.push({score: this.score, name: "$this"});
            }

            if (this.isNewHighScore) {
                this.$nextTick(this.focusPlayerNameInput);
                this.$nextTick(this.addFocusForcerToPlayerNameInput);
            }
        },

        addFocusForcerToPlayerNameInput() {
            if (this.$refs.playerNameInput && this.$refs.playerNameInput[0]) {
                this.$refs.playerNameInput[0].addEventListener("blur", this.focusPlayerNameInput);
            }
        },

        removeFocusForcerFromPlayerNameInput() {
            if (this.$refs.playerNameInput && this.$refs.playerNameInput[0]) {
                this.$refs.playerNameInput[0].removeEventListener("blur", this.focusPlayerNameInput);
            }
        },

        focusPlayerNameInput() {
            if (this.$refs.playerNameInput && this.$refs.playerNameInput[0]) {
                this.$refs.playerNameInput[0].focus();
            }
        },

        async postScore(event) {
            event.preventDefault();
            await fetch("/api/asteroids_scores", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    score: this.score,
                    name: this.playerName,
                })
            });
            this.scoreWasPosted = true;
        }
    }
}
</script>

<style scoped>
span {
    text-align: center;
    display: block;
}

.score-info {
    margin-bottom: .5rem;
}

.game-over-score {
    font-size: 1.25rem;
}

.new-high-score {
    font-size: 1.25rem;
}

.game-over .leaderboard ol {
    list-style-type: none;
    padding: 0;
    margin: 1rem auto;
}

.game-over .leaderboard ol li > span {
    margin-bottom: .2rem;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
}

.game-over .leaderboard ol li span, .game-over .leaderboard ol li input {
    font-size: .8rem;
}

.game-over .leaderboard ol li input {
    text-align: right;
    width: 3rem;
    padding: 0;
    background-color: transparent;
}
</style>