<template>
    <div>
        <header>
            <div class="header-bar">
                <router-link exact to="/" class="site-title" @click.native="closeDrawer">
                    Henry Woody
                </router-link>
                <nav>
                    <ul>
                        <li v-for="link in links" :key="link.title">
                            <router-link
                                :exact="link.exact"
                                :to="link.to"
                                :title="link.title"
                            >
                                {{ link.title }}
                            </router-link>
                        </li>
                    </ul>
                </nav>

                <button class="btn-drawer" @click="toggleDrawer">
                    <span :class="drawerButtonTextClass">&times;</span>
                </button>
                <DarkModeManager/>
            </div>

            <div :class="drawerClass">
                <nav>
                    <ul>
                        <li v-for="link in links" :key="link.title">
                            <router-link
                                :exact="link.exact"
                                :to="link.to"
                                :title="link.title"
                                @click.native="toggleDrawer"
                            >
                                {{ link.title }}
                            </router-link>
                        </li>
                    </ul>
                </nav>

                <DarkModeManager/>
            </div>
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
            drawerIsCollapsed: true,
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
    },
    computed: {
        drawerClass() {
            return this.drawerIsCollapsed ? "header-drawer collapsed" : "header-drawer";
        },
        drawerButtonTextClass() {
            return this.drawerIsCollapsed ? "rotated" : "";
        }
    },
    methods: {
        toggleDrawer() {
            if (this.drawerIsCollapsed) {
                this.openDrawer();
            } else {
                this.closeDrawer();
            }
        },

        openDrawer() {
            this.drawerIsCollapsed = false;
            document.body.classList.add("no-scroll");
            document.querySelector("main").addEventListener("click", this.toggleDrawer);
        },

        closeDrawer() {
            this.drawerIsCollapsed = true;
            document.body.classList.remove("no-scroll");
            document.querySelector("main").removeEventListener("click", this.toggleDrawer);
        },
    }
}
</script>

<style scoped>
header {
    --header-drawer-transition-timing: .3s ease-out;
    --button-height: 1.5rem;
}

.header-bar, .header-drawer {
    background-color: var(--header-background-color);
}

.header-bar, .header-drawer, a {
    color: var(--header-color);

    transition: var(--dark-mode-transition);
}

.header-bar {
    width: 100vw;
    height: var(--header-height);
    padding: 0 1rem;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
    align-items: center;
    box-sizing: border-box;
    position: absolute;
    top: 0;
    left: 0;
    z-index: 1000;
}

.header-drawer {
    --header-drawer-width: 12rem;
    width: var(--header-drawer-width);
    height: 100vh;
    position: fixed;
    left: 0;
    top: 0;
    z-index: 999;

    transition:
        var(--dark-mode-transition),
        left var(--header-drawer-transition-timing);
}

.header-drawer.collapsed {
    left: calc(-1 * var(--header-drawer-width));
}

.header-drawer nav {
    margin-top: 3rem;
    display: flex;
    flex-flow: column;
}

.header-drawer nav li {
    display: block;
    margin-bottom: 1rem;
}

.site-title {
    font-size: 1.25rem;
    font-weight: 600;
}

nav ul {
    list-style-type: none;
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

.header-bar button.btn-drawer {
    width: var(--button-height);
    padding: 0;
    display: flex;
    justify-content: center;
    position: relative;
}

.header-bar button.btn-drawer span {
    --span-height: 1.1rem;
    --span-width: 1.1rem;
    font-size: 1.75rem;
    width: var(--span-width);
    height: var(--span-height);
    line-height: .85rem;
    text-align: center;
    display: block;
    position: absolute;
    left: calc((var(--button-height) - var(--span-width)) / 2);
    top: calc((var(--button-height) - var(--span-height)) / 2);
    transition:
        transform var(--header-drawer-transition-timing);
}

.header-bar button.btn-drawer span.rotated {
    transform: rotate(-45deg);
    transform-origin: center center;
}

.header-drawer button.btn-dark-mode {
    margin: 2rem auto 0;
    display: block;
}

.header-bar nav ul, .header-bar button.btn-dark-mode {
    display: none;
}

button {
    color: var(--color);
    background-color: var(--background-color);
}

@media all and (min-width: 750px) {
    .header-bar nav ul {
        display: flex;
        flex-flow: row;
        justify-content: space-between;
    }

    .header-bar nav li {
        display: inline-block;
    }

    .header-bar nav li:not(:last-of-type) {
        margin-right: 1.5rem;
    }

    .header-bar button.btn-dark-mode {
        display: inline-block;
    }

    .header-bar button.btn-drawer {
        display: none;
    }

    .header-drawer {
        display: none;
    }
}
</style>