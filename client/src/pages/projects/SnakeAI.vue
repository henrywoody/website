<template>
    <main>
        <h1>Snake AI</h1>

        <section class="description">
            <p>
                An artificially intelligent player for Snake. The snake's brain is a neural network and is improved using the genetic algorithm. Written in Python using Pygame for graphics/user interaction. Evolution handled by <router-link to="/projects/holland">Holland</router-link>.
            </p>

            <div class="video-container">
                <video playsinline autoplay>
                    <source src="https://github.com/henrywoody/snake-ai/raw/master/gallery/sample2.mov" type="video/mp4">
                </video>
            </div>

            <p>
                See the Github repository's <ExternalLink href="https://github.com/henrywoody/snake-ai/blob/master/gallery">gallery</ExternalLink> for more sample runs, corresponding genomes are in <ExternalLink href="https://github.com/henrywoody/snake-ai/blob/master/samples">samples</ExternalLink>.
            </p>

            <h2>Snake Anatomy</h2>

            <p>
                Snakes have 5 eyes and can see the closest object that falls along the straight line starting at the snake's head at the angle of the eye.
            </p>

            <p>
                Snakes brains are neural nets with one hidden layer, 15 input nodes, and 5 output nodes. The 15 inputs correspond to the inputs of the eyes. Each eye receives 3 inputs—2 are the visual encoding of the object, and the final is the distance from the snake's head to the object. Each of the 5 output nodes correspond to turning in the direction of one of the snake's eyes (one eye points forward so this is just not turning).
            </p>
        </section>

        <GithubLinkIcon href="https://github.com/henrywoody/snake-ai"/>
    </main>
</template>

<script>
import ExternalLink from "../../components/ExternalLink";
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "SnakeAI",
    components: {
        ExternalLink,
        GithubLinkIcon,
    }
}
</script>

<style scoped>

</style>