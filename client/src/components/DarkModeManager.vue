<template>
    <button class="btn-dark-mode" :class="{ 'dark-on': isDarkModeOn }" @click="toggleDarkMode" title="Dark Mode"></button>
</template>

<script>
export default {
    name: "DarkModeManager",
    data() {
        const isDarkModeOn = window.localStorage.getItem("isDarkModeOn") === "true";

        return {
            isDarkModeOn,
        }
    },
    methods: {
        toggleDarkMode() {
            const body = document.querySelector("body");
            if (body.classList.contains("dark")) {
                body.classList.remove("dark");
                window.localStorage.setItem("isDarkModeOn", false);
                this.isDarkModeOn = false;
            } else {
                body.classList.add("dark");
                window.localStorage.setItem("isDarkModeOn", true);
                this.isDarkModeOn = true;
            }
        },
    }
}
</script>

<style scoped>
button.btn-dark-mode {
    background-color: #777;
    --button-width: 2.2rem;
    width: var(--button-width);
    height: .5rem;
    padding: 0;
    border-radius: .25rem;
    position: relative;
}

button.btn-dark-mode:active {
    transform: none;
}

button.btn-dark-mode::after {
    content: "";
    background-color: var(--background-color);
    --size: 18px;
    width: var(--size);
    height: var(--size);
    border-radius: 50%;
    margin: auto;
    display: inline-block;
    position: absolute;
    left: 0;
    right: 0;
    top: calc(-1 * var(--size) / 4);
    --distance-from-edge: .8rem;
    transform: translateX(calc(-1 * (var(--button-width) / 2 - var(--distance-from-edge))));
    transition: transform .1s ease-out, color calc(var(--dark-mode-transition-time) - .05s) linear .05s;
}

button.btn-dark-mode.dark-on::after {
    transform: translateX(calc(var(--button-width) / 2 - var(--distance-from-edge)));
}
</style>