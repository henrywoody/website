<template>
    <div v-on:click="this.goToPage" class="project-item">
        <span class="title">{{ title }}</span>

        <p>
            {{ description }}
        </p>
    </div>
</template>

<script>
export default {
    name: "ProjectItem",
    props: {
        title: String,
        description: String,
        path: String
    },
    methods: {
        goToPage() {
            this.$router.push(this.path);
        }
    }
}
</script>

<style scoped>
.project-item {
    padding: 1rem;
    margin-bottom: 1.5rem;

    transition: .3s;
}

.project-item:hover {
    box-shadow: var(--box-shadow);
    transform: var(--box-shadow-translate);
    cursor: pointer;

    transition: .15s;
}

.project-item:active {
    box-shadow: var(--box-shadow-inset);
    transform: var(--box-shadow-inset-translate);    

    transition: .1s;
}

.title {
    font-size: 1.25rem;
    font-weight: 300;
}

p {
    margin: 1rem 2rem;
}
</style>