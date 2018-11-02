<template>
    <main>
        <h1>Drobe</h1>

        <span class="link"><a href="https://drobe.io">drobe.io</a></span>

        <p>
            Drobe is an application designed to save time and brain power in the morning by picking an outfit for a user to wear. After creating an account, a user can create records of the clothes in their wardrobe.
        </p>

        <p>
            The most important feature for creating outfits is the related articles of clothing for each article. These are speficied by the user. Essentially the user decides which pants go with which shirts and so on.
        </p>

        <p>
            Each article in an outfit is selected probabilistically according to:
            <ul>
                <li>The user's preference for that article</li>
                <li>The weather: temperature and rain</li>
                <li>The last time the article was worn</li>
                <li>The other articles so far selected for the outfit</li>
            </ul>
        </p>

        <p>
            Weather data comes from <a href="https://darksky.net/poweredby/">Dark Sky's Weather API</a>. This API gives weather data for a given set of coordinates. Coordinate data comes from <a href="https://developers.google.com/maps/documentation/geocoding/start">Google's Geocoding API</a>.
        </p>

        <p>
            The Drobe server was written in Node using the Express framework. The frontend is in React. The database is Postgres. Drobe is currently hosted on Google Cloud's App Engine.
        </p>

        <GithubLinkIcon href="https://github.com/henrywoody/drobe"/>
    </main>
</template>

<script>
import GithubLinkIcon from '../../components/GithubLinkIcon.vue';

export default {
    name: "Drobe",
    components: {
        GithubLinkIcon
    }
}
</script>

<style scoped>
.link {
    font-size: 1.1rem;
    text-align: center;
    margin: 0 auto 1rem;
    display: block;
}

.link a:hover {
    color: #1bb8d1; /*drobe blue*/
}

p {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}
</style>