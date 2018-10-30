<template>
    <main>
        <h1>Projects</h1>

        <p>
            Projects are my life. Here are a few of my recent projects. If you're trying to work on something neat, hit me up.
        </p>

        <section
            v-for="projectSection in projectSections"
            v-bind:key="projectSection.title"
            class="projects-list"
        >
            <h2>{{ projectSection.title }}</h2>

            <ProjectItem
                v-for="project in projectSection.projects"
                v-bind:key="project.title"
                v-bind:title="project.title"
                v-bind:description="project.description"
                v-bind:path="project.path"
            />
        </section>

        <p>
            Check out my <a href="https://github.com/henrywoody">GitHub page</a> for more!
        </p>
    </main>
</template>

<script>
import ProjectItem from '../components/ProjectItem.vue';

export default {
    name: "Projects",
    components: {
        ProjectItem
    },
    data() {
        return {
            projectSections: [
                {
                    title: "Web Development",
                    projects: [
                        {
                            title: "Drobe",
                            description: "Programmatically curated outfits",
                            path: "/projects/drobe"
                        },
                        {
                            title: "This Site",
                            description: "$CURR_LOCATION",
                            path: "/projects/website"
                        }
                    ]
                },
                {
                    title: "Artificial Life",
                    projects: [
                        {
                            title: "One Dimensional Cellular Automata",
                            description: "Life at its simplest",
                            path: "/projects/1d-cellular-automata"
                        },
                        {
                            title: "Two Dimensional Cellular Automata",
                            description: "Things get a bit more complex",
                            path: "/projects/2d-cellular-automata"
                        },
                        {
                            title: "Boids",
                            description: "Birds of a feather...",
                            path: "/projects/boids"
                        },
                        {
                            title: "Holland",
                            description: "A genetic algorithm for any application",
                            path: "/projects/holland"
                        }
                    ]
                }
            ]
        }
    }
}
</script>