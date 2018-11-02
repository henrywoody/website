<template>
    <div id="app">
        <Header/>
        <router-view/>
    </div>
</template>

<script>
import Header from './components/Header.vue';

export default {
    name: 'app',
    components: {
        Header
    },
    created() {
        const isDarkModeOn = window.localStorage.getItem("isDarkModeOn") === "true";
        if (isDarkModeOn) {
            const body = document.querySelector("body");
            body.classList.add("dark");
        }
    }
}
</script>

<style>
body {
    --color: #333;
    --inverted-color: #eee;
    --background-color: #fff;
    --inverted-background-color: #333;
    --shadow-color: #777;
    --header-color: var(--inverted-color);
    --header-background-color: var(--inverted-background-color);

    --shadow-offset: .1rem;
    --shadow-blur: .2rem;
    --box-shadow:
        var(--shadow-offset) var(--shadow-offset) var(--shadow-blur) var(--shadow-color),
        0 0 var(--shadow-blur) var(--shadow-color);
    --box-shadow-translate: translate(0, calc(-1 * var(--shadow-offset)));
    --box-shadow-inset:
        inset 0 0 var(--shadow-blur) var(--shadow-color),
        inset var(--shadow-offset) var(--shadow-offset) var(--shadow-blur) var(--shadow-color);
    --box-shadow-inset-translate: translate(0, var(--shadow-offset));

    --header-height: 2.5rem;

    font-family: "Open Sans", sans-serif, Arial;
    font-weight: 400;
    color: var(--color);
    background-color: var(--background-color);
    margin: 0;

    --dark-mode-transition-time: .3s;
    --dark-mode-transition:
        var(--dark-mode-transition-time) color,
        var(--dark-mode-transition-time) background-color;
    transition: var(--dark-mode-transition);
}

body.dark {
    --color: #eee;
    --inverted-color: #333;
    --background-color: #333;
    --inverted-background-color: #fff;
    --shadow-color: #000;
    --header-color: var(--color);
    --header-background-color: #555;
}

main {
    padding: 0 2rem;
}

h1 {
    font-size: 2.5rem;
    font-weight: 300;
}

h2 {
    font-weight: 2rem;
    font-weight: 300;
}

a {
    color: var(--color);
    transition: var(--dark-mode-transition);
}

button {
        outline: none;
        border: none;
}
</style>