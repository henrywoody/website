<template>
    <router-link :to="path" class="project-item">
        <span class="title">{{ title }}</span>

        <p>
            {{ description }}
        </p>
    </router-link>
</template>

<script>
export default {
    name: "ProjectItem",
    props: {
        title: String,
        description: String,
        path: String
    },
}
</script>

<style scoped>
.project-item {
    text-decoration: none;
    padding: 1rem 1rem 1.5rem;
    margin-bottom: 1rem;
    display: block;

    --transition-timing: .3s;
    transition: var(--transition-timing) box-shadow, var(--transition-timing) transform;
}

.project-item:hover {
    box-shadow: var(--box-shadow);
    transform: var(--box-shadow-translate);
    cursor: pointer;

    --transition-timing: .15s;
}

.project-item:active {
    box-shadow: var(--box-shadow-inset);
    transform: var(--box-shadow-inset-translate);    

    transition: .1s;
}

.title {
    font-size: 1.25rem;
    font-weight: 400;
}

p {
    margin: 1rem 2rem 0;
    font-size: .9rem;
}
</style>