<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>
export default {
    name: 'app'
}
</script>

<style>
body {
    --color: #222;
    --background-color: #fff;
    --shadow-color: #222;

    --shadow-offset: .1rem;
    --shadow-blur: .2rem;
    --box-shadow:
        var(--shadow-offset) var(--shadow-offset) var(--shadow-blur) var(--shadow-color),
        0 0 var(--shadow-blur) var(--shadow-color);
    --box-shadow-translate: translate(calc(-1 * var(--shadow-offset)), calc(-1 * var(--shadow-offset)));
    --box-shadow-inset:
        inset 0 0 var(--shadow-blur) var(--shadow-color),
        inset var(--shadow-offset) var(--shadow-offset) var(--shadow-blur) var(--shadow-color);
    --box-shadow-inset-translate: translate(var(--shadow-offset), var(--shadow-offset));

    font-family: "Open Sans", sans-serif, Arial;
    font-weight: 400;
    color: var(--color);
    background-color: var(--background-color);
    margin: 0;
}

body.dark {
    --color: #eee;
    --background-color: #222;
    --shadow-color: #888;
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
}
</style>