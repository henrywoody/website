<template>
    <main>
        <h1>Drobe</h1>

        <div class="link-container">
            <ExternalLink href="https://drobe.io">drobe.io</ExternalLink>
        </div>

        <section class="description">
            <p>
                Drobe is an application that picks outfits for users to wear each day. After creating an account, a user can create records of the clothes in their wardrobe. The app was designed to save time and brain power in the morning.
            </p>

            <p>

                The most important feature for creating outfits is the related articles of clothing for each article. These are specified by the user. The user decides, for example, which pairs of pants go with which shirts.
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
                Weather data comes from <ExternalLink href="https://darksky.net/poweredby/">Dark Sky's Weather API</ExternalLink>. This API gives weather data for a given set of coordinates. Coordinate data comes from <ExternalLink href="https://developers.google.com/maps/documentation/geocoding/start">Google's Geocoding API</ExternalLink>.
            </p>

            <p>
                The Drobe server was written in Node using the Express framework. The frontend is in React. The database was originally written in MongoDB, but due to to the (many-to-many) relationships between the data, I recreated the database using Postgres. Drobe is currently hosted on Google Cloud's App Engine.
            </p>
        </section>

        <GithubLinkIcon href="https://github.com/henrywoody/drobe"/>
    </main>
</template>

<script>
import ExternalLink from "../../components/ExternalLink";
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "Drobe",
    components: {
        ExternalLink,
        GithubLinkIcon,
    }
}
</script>

<style scoped>
.link-container {
    font-size: 1.1rem;
    text-align: center;
    margin: 0 auto 1rem;
}

.link-container a:hover {
    color: #1bb8d1; /* drobe blue */
}

p {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}
</style>