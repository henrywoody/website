<template>
    <main>
        <h1>KandinskyBot</h1>

        <section class="description">
            <div class="quote-container">
                <blockquote>
                    Everything starts from a dot.
                </blockquote>

                <span class="quote-attribution">Wassily Kandinsky</span>
            </div>

            <p>
                This program creates random works of "art" in the style of Wassily Kandinsky. Mostly inspired by <em>Composition VIII</em>, though <em>Color Study. Squares with Concentric Circles</em> was an influence as well.
            </p>

            <div class="image-container">
                <img src="https://camo.githubusercontent.com/bad186dcd807c197afbc3f65bfdf3d80b7d12305/68747470733a2f2f7777772e77617373696c796b616e64696e736b792e6e65742f696d616765732f776f726b732f35302e6a7067" alt="Composition VIII">
                <span class="image-description">Kandinsky, W. (1923). Composition VIII [Painting]. <ExternalLink href="https://www.wassilykandinsky.net">Source</ExternalLink>.</span>
            </div>

            <p>
                While most of the images produced by KandinskyBot are trash, there is the odd gem—around 1-2%. The infrequency of quality images is part of the fun.
            </p>

            <p>
                Here are a few samples:
            </p>

            <div class="image-container">
                <img src="https://raw.githubusercontent.com/henrywoody/kandinsky-bot/master/samples/1141.png" alt="Composition 1141">
                <span class="image-description">KandinskyBot (2017). Composition 1141 [Digital].</span>
            </div>

            <div class="image-container">
                <img src="https://raw.githubusercontent.com/henrywoody/kandinsky-bot/master/samples/113.png" alt="Composition 113">
                <span class="image-description">KandinskyBot (2017). Composition 113 [Digital].</span>
            </div>

            <div class="image-container">
                <img src="https://raw.githubusercontent.com/henrywoody/kandinsky-bot/master/samples/534.png" alt="Composition 534">
                <span class="image-description">KandinskyBot (2017). Composition 534 [Digital].</span>
            </div>

            <p>
                More samples in the <ExternalLink href="https://github.com/henrywoody/kandinsky-bot/tree/master/samples">samples directory</ExternalLink> on Github.
            </p>

            <p>
                Kandinsky bot also has an Instagram account that shows images it's generated. Check it out: <ExternalLink href="https://www.instagram.com/kandinskybot">@KandinskyBot</ExternalLink>.
            </p>
        </section>

        <GithubLinkIcon href="https://github.com/henrywoody/kandinsky-bot"/>
    </main>
</template>

<script>
import ExternalLink from "../../components/ExternalLink";
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "KandinskyBot",
    components: {
        ExternalLink,
        GithubLinkIcon,
    }
}
</script>
