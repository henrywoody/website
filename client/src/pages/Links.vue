<template>
    <main>
        <h1>Links</h1>

        <p>
            Some great sites.
        </p>

        <section v-for="section in sections" v-bind:key="section.title">
            <h2>{{ section.title }}</h2>

            <ul>
                <li v-for="link in section.links" v-bind:key="link.title">
                    <a v-bind:href="link.href">{{ link.title }}</a>
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
                    }
                ]
            }
        ].map(e => ({
            ...e,
            links: e.links.sort((a,b) => {
                return a.title.split(" ")[1].charCodeAt(0) - b.title.split(" ")[1].charCodeAt(0)
            })
        }));
        
        return { sections }
    }
}
</script>