<template>
    <main>
        <h1>Projects</h1>

        <p>
            Projects are my life.
        </p>

        <section
            v-for="projectSection in projectSections"
            :key="projectSection.title"
            class="projects-list"
        >
            <h2>{{ projectSection.title }}</h2>

            <ProjectItem
                v-for="project in projectSection.projects"
                :key="project.title"
                :title="project.title"
                :description="project.description"
                :path="project.path"
            />
        </section>

        <p>
            Check out my <ExternalLink href="https://github.com/henrywoody">GitHub page</ExternalLink> for more!
        </p>
    </main>
</template>

<script>
import ProjectItem from "../components/ProjectItem";
import ExternalLink from "../components/ExternalLink";

export default {
    name: "Projects",
    components: {
        ProjectItem,
        ExternalLink,
    },
    data() {
        return {
            projectSections: [
                {
                    title: "Web Development",
                    projects: [
                        {
                            title: "Drobe",
                            description: "Programmatically curated outfits.",
                            path: "/projects/drobe",
                        },
                        {
                            title: "This Site",
                            description: "$CURR_LOCATION",
                            path: "/projects/website",
                        },
                        {
                            title: "404",
                            description: "Feeling lost in cyberspace?",
                            path: "/404",
                        }
                    ]
                },
                {
                    title: "Artificial Life",
                    projects: [
                        {
                            title: "Holland",
                            description: "A genetic algorithm for any application.",
                            path: "/projects/holland",
                        },
                        {
                            title: "Boids",
                            description: "Birds of a feather...",
                            path: "/projects/boids",
                        },
                        {
                            title: "One Dimensional Cellular Automata",
                            description: "Life at its simplest.",
                            path: "/projects/1d-cellular-automata",
                        },
                        {
                            title: "Two Dimensional Cellular Automata",
                            description: "Things get a bit more complex.",
                            path: "/projects/2d-cellular-automata",
                        },
                        {
                            title: "Snake AI",
                            description: "A mind of its own.",
                            path: "/projects/snake-ai",
                        }
                    ]
                },
                {
                    title: "Art",
                    projects: [
                        {
                            title: "KandinskyBot",
                            description: "The world's most advanced artist.",
                            path: "/projects/kandinsky-bot",
                        },
                        {
                            title: "Barnsley Fern",
                            description: "The fractal fern.",
                            path: "/projects/barnsley-fern",
                        }
                    ]
                },
                {
                    title: "CSS Projects",
                    projects: [
                        {
                            title: "CSS Flags",
                            description: "The world's flags recreated in CSS.",
                            path: "/projects/css-flags",
                        }
                    ],
                },
            ]
        }
    }
}
</script>