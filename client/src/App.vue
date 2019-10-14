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
    --accent-background-color: #eee;
    --inverted-background-color: #333;
    --shadow-color: #777;
    --header-color: var(--inverted-color);
    --header-background-color: var(--inverted-background-color);
    --button-background-color: var(--inverted-color);

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

    --font-family: "Open Sans", sans-serif, Arial;
    font-family: var(--font-family);
    font-weight: 400;
    color: var(--color);
    background-color: var(--background-color);
    margin: 0;
    overflow: auto;

    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;

    --dark-mode-transition-time: .3s;
    --dark-mode-transition:
        color var(--dark-mode-transition-time),
        background-color var(--dark-mode-transition-time),
        fill var(--dark-mode-transition-time);
    transition: var(--dark-mode-transition);
}

body.dark {
    --color: #eee;
    --inverted-color: #333;
    --background-color: #333;
    --accent-background-color: #555;
    --inverted-background-color: #fff;
    --shadow-color: #000;
    --header-color: var(--color);
    --header-background-color: #555;
    --button-background-color: #555;
}

main {
    padding: 0 2rem 1rem;
    margin-top: var(--header-height);
    box-sizing: border-box;
    overflow: auto;
}

h1 {
    font-size: 2.5rem;
    font-weight: 400;
}

h2 {
    font-weight: 2rem;
    font-weight: 400;
}

a {
    color: var(--color);
    transition: var(--dark-mode-transition);
}

input {
    color: var(--color);
    font-size: 1.2rem;
    font-family: var(--font-family);
    background-color: var(--background-color);
    width: 5rem;
    outline: none;
    border: none;

    transition: var(--dark-mode-transition);
}

button {
    color: var(--color);
    background-color: var(--button-background-color);
    height: 1.5rem;
    outline: none;
    border: none;

    transition: var(--dark-mode-transition);
}

button:active {
    transform: translate(.05rem, .05rem);
}

.fixed {
    position: fixed !important;
}

.no-scroll {
    overflow: hidden;
}

section.description {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}


input[type="range"] {
    display: none;
}

@media all and (min-width: 800px) {
    input[type="range"] {
        display: initial;
    }

    input.range-standin {
        display: none;
    }
}

@media print {
    header {
        display: none !important;
    }

    main {
        margin-top: 0;
    }
}

.quote-container {
    width: fit-content;
    margin: 0 auto 2rem;
    display: flex;
    flex-flow: column;
    align-items: flex-end;
}

.quote-container blockquote {
    background-color: var(--accent-background-color);
    width: fit-content;
    padding: .8rem 1.2rem;
    margin: 0 0 .5rem;
    transition: background-color var(--dark-mode-transition-time);
}

.quote-container .quote-attribution {
    width: fit-content;
    padding-right: .5rem;
}

.quote-container .quote-attribution::before {
    content: "―";
}

.image-container {
    width: 100%;
    margin-bottom: 1rem;
}

.image-container img {
    width: 100%;
}

.image-container .image-description {
    font-size: .9rem;
}

.video-container {
    width: 100%;
}

.video-container video {
    width: 100%;
}
</style>