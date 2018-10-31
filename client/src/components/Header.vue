<template>
    <div>
        <header>
            <router-link exact to="/" class="site-title">Henry Woody</router-link>
            <nav>
                <ul>
                    <li v-for="link in links" v-bind:key="link.title">
                        <router-link
                            v-bind:exact="link.exact"
                            v-bind:to="link.to"
                            v-bind:title="link.title"
                        >
                            {{ link.title }}
                        </router-link>
                    </li>
                </ul>
            </nav>

            <DarkModeManager/>
        </header>
    </div>
</template>

<script>
import DarkModeManager from './DarkModeManager.vue';

export default {
    name: "Header",
    components: {
        DarkModeManager
    },
    data() {
        return {
            links: [
                {
                    title: "Home",
                    to: "/",
                    exact: true
                },
                {
                    title: "Projects",
                    to: "/projects",
                    exact: false
                },
                {
                    title: "Links",
                    to: "/links",
                    exact: false
                }
            ]
        }
    }
}
</script>

<style scoped>
header {
    background-color: var(--inverted-background-color);
    height: var(--header-height);
    padding: 0 1rem;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
    align-items: center;
}

header, a {
    color: var(--inverted-color);

    transition: var(--dark-mode-transition);
}

.site-title {
    font-size: 1.25rem;
    font-weight: 400;
}

nav ul {
    list-style-type: none;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
}

nav li {
    display: inline-block;
}
nav li:not(:last-of-type) {
    margin-right: 1.5rem;
}

a {
    text-decoration: none;
    letter-spacing: .05rem;
}

a::after {
    content: attr(title);
    font-weight: 600;
    height: 0;
    display: block;
    overflow: hidden;
    visibility: hidden;
}

a.router-link-active:not(.site-title) {
    font-weight: 600;
}

a.router-link-exact-active {
    cursor: default;
}

button {
    color: var(--color);
    background-color: var(--background-color);
}
</style>