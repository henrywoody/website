<template>
    <button class="btn-dark-mode" @click="toggleDarkMode">{{ isDarkModeOn ? "Light" : "Dark" }}</button>
</template>

<script>
export default {
    name: "DarkModeManager",
    data() {
        const isDarkModeOn = window.localStorage.getItem("isDarkModeOn") === "true";

        return {
            isDarkModeOn
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
button {
    width: 3rem;
    height: 1.5rem;
}
</style>