<template>
    <main>
        <h1>Links</h1>

        <p>
            Some great sites.
        </p>

        <section v-for="section in sections" :key="section.title">
            <h2>{{ section.title }}</h2>

            <ul>
                <li v-for="link in section.links" :key="link.title">
                    <a :href="link.href">{{ link.title }}</a>
                </li>
            </ul>
        </section>
    </main>
</template>

<script>
export default {
    name: "Links",
    data() {
        const sections = [
            {
                title: "Web",
                links: [
                    {
                        title: "Future Retro Redesigns",
                        href: "https://interface.fh-potsdam.de/future-retro/"
                    },
                    {
                        title: "7UP Netherlands",
                        href: "https://www.7up.nl/7up"
                    }
                ]
            },
            {
                title: "Artificial Life & Intelligence",
                links: [
                    {
                        title: "Primordial Soup Kitchen",
                        href: "http://psoup.math.wisc.edu/welcome.html"
                    },
                    {
                        title: "Genetic Algorithms & Artificial Life YouTube Playlist",
                        href: "https://www.youtube.com/playlist?list=PLhhW7H3Rz9W1Z59WwWwxEMofNyUc-0jBb"
                    }
                ]
            },
            {
                title: "Math",
                links: [
                    {
                        title: "Animated Factorization",
                        href: "http://www.datapointed.net/visualizations/math/factorization/animated-diagrams/"
                    }
                ]
            },
            {
                title: "Friends",
                links: [
                    {
                        title: "Andrew Langtry",
                        href: "https://andrewlangtry.com"
                    },
                    {
                        title: "Jesse Jenks",
                        href: "https://www.verychill.biz"
                    },
                    {
                        title: "Banji Oyewole",
                        href: "http://banj.io"
                    },
                    {
                        title: "Chili Johnson",
                        href: "https://chilij.com"
                    },
                    {
                        title: "Ethan Russell",
                        href: "http://ethan-russell.com"
                    },
                    {
                        title: "Joe Marin",
                        href: "https://joemarin.com/"
                    }
                ].sort((a, b) => {
                    return a.title.split(" ")[1].charCodeAt(0) - b.title.split(" ")[1].charCodeAt(0);
                })
            }
        ]
        
        return { sections }
    }
}
</script>