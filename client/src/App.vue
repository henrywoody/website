<template>
    <div id="app">
        <Index/>
    </div>
</template>

<script>
import Index from './pages/Index.vue';

export default {
    name: 'app',
    components: {
        Index
    }
}
</script>

<style>
body {
    font-family: "Open Sans", Arial;
    font-weight: 300;
    background-color: #111;
    margin: 0;
}
</style>