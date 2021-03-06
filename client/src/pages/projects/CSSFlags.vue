<template>
    <main>
        <h1>CSS Flags</h1>

        <div class="link-container">
            <ExternalLink href="http://flags.henrywoody.com">flags.henrywoody.com</ExternalLink>
        </div>

        <div class="flag-container">
            <div class="flag-css-flags"/>
        </div>

        <section class="description">
            <p>
                To practice my CSS skills and learn more about flags and their countries, I decided to recreate real-world flags using only HTML and CSS.
            </p>

            <p>
                The first flag I recreated was that of Japan. I was learning about CSS gradients,
                specifically the <code>radial-gradient</code>, and learned about hard color stops,
                which let you change color in a gradient with a hard line rather than a blend like usual.
                From there I set out to put those skills to the test.
            </p>

            <p>
                All flags in the project were created with a single <code>div</code> element and using only CSS for each design.
            </p>

            <p>
                This project has taught me more than I thought I could ever know about CSS gradients,
                refreshed my knowledge of trigonometry, let me learn more about the flags of many different nations,
                and shown me what makes good and bad flag design.
                Most of the work on each flag is translating the flag construction sheet (if there is one) into positions and angles for stripes and triangles,
                which can be translated into CSS gradient rules.
            </p>

            <p>
                The image at the top of the page is the flag I designed for the project itself.
                Its design combines the flags for the letters of CSS using the <ExternalLink href="https://en.wikipedia.org/wiki/International_Code_of_Signals">International Code of Signals</ExternalLink> with the two S flags overlayed on top of the C flag.
            </p>
        </section>

        <GithubLinkIcon href="https://github.com/henrywoody/css-flags"/>
    </main>
</template>

<script>
import ExternalLink from "../../components/ExternalLink";
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "CSSFlags",
    components: {
        ExternalLink,
        GithubLinkIcon,
    },
}
</script>

<style scoped>
.link-container {
    font-size: 1.1rem;
    text-align: center;
    margin: 0 auto 1rem;
}

.flag-container {
    width: 100%;
    margin-bottom: 3rem;
}

.flag-css-flags {
    --blue: #368ce3;
    --white: #fff;
    --red: #b5130d;

    --aspect-ratio: calc(1 / 2);

    --width: 75vw;
    --height: calc(var(--aspect-ratio) * var(--width));

    --fimbriation-width: calc(1 / 54 * var(--width));

    width: var(--width);
    height: var(--height);
    margin: 0 auto;

    background:
        /* S */
        linear-gradient(
            0deg,
            var(--blue) 0 100%
        ) no-repeat calc(1 / 6 * var(--width)) calc(1 / 3 * var(--height)) / calc(1 / 6 * var(--width)) calc(1 / 3 * var(--height)),
        linear-gradient(
            0deg,
            var(--white) 0 100%
        ) no-repeat calc(1 / 6 * var(--width) - var(--fimbriation-width)) calc(1 / 3 * var(--height) - var(--fimbriation-width)) / calc(1 / 6 * var(--width) + 2 * var(--fimbriation-width)) calc(1 / 3 * var(--height) + 2 * var(--fimbriation-width)),

        /* S */
        linear-gradient(
            0deg,
            var(--blue) 0 100%
        ) no-repeat calc(4 / 6 * var(--width)) calc(1 / 3 * var(--height)) / calc(1 / 6 * var(--width)) calc(1 / 3 * var(--height)),
        linear-gradient(
            0deg,
            var(--white) 0 100%
        ) no-repeat calc(4 / 6 * var(--width) - var(--fimbriation-width)) calc(1 / 3 * var(--height) - var(--fimbriation-width)) / calc(1 / 6 * var(--width) + 2 * var(--fimbriation-width)) calc(1 / 3 * var(--height) + 2 * var(--fimbriation-width)),

        /* C */
        linear-gradient(
            180deg,
            var(--blue) 0 calc(1 / 5 * var(--height)),
            var(--white) 0 calc(2 / 5 * var(--height)),
            var(--red) 0 calc(3 / 5 * var(--height)),
            var(--white) 0 calc(4 / 5 * var(--height)),
            var(--blue) 0 calc(5 / 5 * var(--height))
        )
    ;
}
</style>