<template>
    <main>
        <h1>404: Not Found</h1>

        <Asteroids/>
    </main>
</template>

<script>
import Asteroids from './projects/Asteroids.vue';

export default {
    name: "NotFound",
    components: {
        Asteroids
    }
}
</script>

<style scoped>
</style>