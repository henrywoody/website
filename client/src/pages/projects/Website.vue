<template>
    <div class="container">
        <div class="top-slide">
            <div>
                <h1>This Site</h1>

                <p>
                    Scroll down to see behind the scenes
                </p>
            </div>
        </div>

        <main class="no-scroll">
            <h1>Website</h1>

            <section>
                <span class="version">Version <code>3.0</code></span>

                <p>
                    This site is where I share projects I've worked on and try out new ideas in web development. Mostly it's for canvas animations. The ability to share your work and ideas on the Internet is what initially drew me to web development.
                </p>

                <p>
                    The frontend is written in Vue, and is the first project I've created in Vue. The server is written in Go, and, again, is my first Go project. The server is hosted at my house on my Raspberry Pi 3.
                </p>

                <GithubLinkIcon href="https://github.com/henrywoody/website"/>
            </section>

            <section>
                <span class="version">Version <code>2.0</code></span>

                <p>
                    The previous version of this site had the same focus and general style as the current version. The main difference is the technology used. Also sometimes you just need to start fresh to keep things clean.
                </p>

                <p>
                    Version 2 of the site was written using the template engine Pug and was served by a Node/Express server.
                </p>

                <GithubLinkIcon href="https://github.com/henrywoody/personal-web"/>
            </section>

            <section>
                <span class="version">Version <code>1.0</code></span>

                <p>
                    Using <code>1.0</code>, instead of <code>0.0</code>, is generious for this particular version of the site. Version 1 was pretty much just a welcome page with a background image. Written, as you might imagine, in plain HTML and CSS.
                </p>

            </section>
        </main>
    </div>
</template>

<script>
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "Website",
    components: {
        GithubLinkIcon,
    },
    mounted() {
        window.scrollTo(0,0);
        const slider = document.querySelector(".top-slide");
        const main = document.querySelector("main");
        this.intervalId = window.setInterval(() => {
            const { y, height } = slider.getBoundingClientRect();
            const bottomMargin = 10; // see below for the 10px in .top-slide margin-bottom
            if (height + bottomMargin + y === 0 && main.classList.contains("no-scroll")) {
                main.classList.remove("no-scroll");
            } else if (height + bottomMargin + y > 0 && !main.classList.contains("no-scroll")) {
                main.classList.add("no-scroll");
            }
        }, 500);
    },
    beforeDestroy() {
        window.clearInterval(this.intervalId);
    }
}
</script>

<style scoped>
.container {
    --slider-height: 100vh;
    --main-height: 100vh;
    height: calc(var(--slider-height) + var(--main-height) + 10px); /* the 10px is used abov in mounted() hook */
}

.top-slide {
    background-color: var(--background-color);
    box-shadow: var(--box-shadow);
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    z-index: 500;

    transition: background-color var(--dark-mode-transition-time);
}

.top-slide > div {
    margin: 0 auto;
}

.top-slide h1, .top-slide p {
    text-align: center;
}

main {
    width: 100vw;
    height: 100vh;
    margin-top: 0;
    position: fixed;
    top: 0;
    left: 0;
}

section {
    margin-bottom: 2rem;
}

.version {
    text-align: center;
    display: block;
    margin-bottom: 1rem;
}

section p {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}

section p:last-of-type {
    margin-bottom: .75rem;
}

.github-icon {
    margin-top: 0;
}
</style>