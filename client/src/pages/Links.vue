<template>
    <main>
        <h1>Links</h1>

        <p>
            Some great sites.
        </p>

        <section>
            <h2>Programming</h2>

            <ul>
                <li>
                    <ExternalLink href="https://www.codewars.com">Codewars</ExternalLink> (<ExternalLink href="https://www.codewars.com/users/henrywoody">my profile</ExternalLink>)
                </li>

                <li>
                    <ExternalLink href="https://css-challenges.com">CSS Challenges</ExternalLink> (<ExternalLink href="https://css-challenges.com/members/henry">my profile</ExternalLink>)
                </li>

                <li>
                    <ExternalLink href="https://github.com">GitHub</ExternalLink> (<ExternalLink href="https://github.com/henrywoody">my profile</ExternalLink>)
                </li>

                <li>
                    <ExternalLink href="https://foobar.withgoogle.com">Google FooBar</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="http://www.pythonchallenge.com">Python Challenge</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://regexcrossword.com">Regex Crossword</ExternalLink> (<ExternalLink href="https://regexcrossword.com/profile/54291">my profile</ExternalLink>)
                </li>

                <li>
                    <ExternalLink href="https://stackoverflow.com">Stack Overflow</ExternalLink> (<ExternalLink href="https://stackoverflow.com/users/8068625/henry-woody?tab=profile">my profile</ExternalLink>)
                </li>
            </ul>
        </section>

        <section>
            <h2>Web</h2>

            <ul>
                <li>
                    <ExternalLink href="https://interface.fh-potsdam.de/future-retro">Future Retro Redesigns</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://www.7up.nl/7up">7UP Netherlands</ExternalLink>
                </li>
            </ul>
        </section>

        <section>
            <h2>Artificial Life &amp; Intelligence</h2>

            <ul>
                <li>
                    <ExternalLink href="http://psoup.math.wisc.edu/welcome.html">Primordial Soup Kitchen</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://www.youtube.com/playlist?list=PLhhW7H3Rz9W1Z59WwWwxEMofNyUc-0jBb">Genetic Algorithms & Artificial Life YouTube Playlist</ExternalLink>
                </li>
            </ul>
        </section>

        <section>
            <h2>Math</h2>

            <ul>
                <li>
                    <ExternalLink href="http://www.datapointed.net/visualizations/math/factorization/animated-diagrams">Animated Factorization</ExternalLink>
                </li>
            </ul>
        </section>

        <section>
            <h2>Friends</h2>

            <ul>
                <li>
                    <ExternalLink href="https://andrewlangtry.com">Andrew Langtry</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://www.verychill.biz">Jesse Jenks</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="http://banj.io">Banji Oyewole</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://chilij.com">Chili Johnson</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="http://ethan-russell.com">Ethan Russell</ExternalLink>
                </li>

                <li>
                    <ExternalLink href="https://joemarin.com/">Joe Marin</ExternalLink>
                </li>
            </ul>
        </section>
    </main>
</template>

<script>
import ExternalLink from "../components/ExternalLink";

export default {
    name: "Links",
    components: {
        ExternalLink,
    },
}
</script>