<template>
    <main>
        <div class="col-container">
            <section class="header">
                <h1>Henry Woody</h1>
            </section>

            <section class="profiles">
                <ul>
                    <li><i><img src="https://www.shareicon.net/data/128x128/2015/09/24/106406_email_512x512.png"></i> <a href="mailto:henry@henrywoody.com">henry@henrywoody.com</a></li>
                    <li><i><img src="https://github.com/favicon.ico"></i> <a href="https://github.com/henrywoody">github.com/henrywoody</a></li>
                    <li><i><img src="https://stackoverflow.com/favicon.ico"></i> <a href="https://stackoverflow.com/story/henrywoody">stackoverflow.com/story/henrywoody</a></li>
                    <li><i><img src="https://codewars.com/favicon.ico"></i> <a href="https://codewars.com/users/henrywoody">codewars.com/users/henrywoody</a></li>
                </ul>
            </section>
        </div>

        <section class="education">
            <h2>Education</h2>

            <div class="res-item">
                <span class="title">BS in Mathematics &amp; BS in Economics</span>
                <span class="organization">University of Puget Sound</span>
                <span class="dates">May 2017</span>

                <ul>
                    <li>Magna Cum Laude</li>
                    <li>|</li>
                    <li>Phi Beta Kappa</li>
                    <li>|</li>
                    <li>Honors in Mathematics</li>
                    <li>|</li>
                    <li>Honors in Economics</li>
                </ul>

            </div>
        </section>

        <section class="skills">
            <h2>Skills</h2>

            <ul>
                <li>
                    <i><img src="https://www.python.org/static/favicon.ico"></i> Python
                </li>
                <li>
                    <i><img src="https://upload.wikimedia.org/wikipedia/commons/6/6a/JavaScript-logo.png"></i> JavaScript
                </li>
                <li>
                    <i><img src="https://golang.org/favicon.ico"></i> Go
                </li>
                <li>
                    <i><img src="https://cdn2.iconfinder.com/data/icons/web-and-mobile-ui-volume-7/48/309-512.png"></i> SQL
                </li>
                <li>
                    <i><img src="https://www.djangoproject.com/favicon.ico"></i> Django
                </li>
                <li>
                    <i><img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9Ii0xMS41IC0xMC4yMzE3NCAyMyAyMC40NjM0OCI+CiAgPHRpdGxlPlJlYWN0IExvZ288L3RpdGxlPgogIDxjaXJjbGUgY3g9IjAiIGN5PSIwIiByPSIyLjA1IiBmaWxsPSIjNjFkYWZiIi8+CiAgPGcgc3Ryb2tlPSIjNjFkYWZiIiBzdHJva2Utd2lkdGg9IjEiIGZpbGw9Im5vbmUiPgogICAgPGVsbGlwc2Ugcng9IjExIiByeT0iNC4yIi8+CiAgICA8ZWxsaXBzZSByeD0iMTEiIHJ5PSI0LjIiIHRyYW5zZm9ybT0icm90YXRlKDYwKSIvPgogICAgPGVsbGlwc2Ugcng9IjExIiByeT0iNC4yIiB0cmFuc2Zvcm09InJvdGF0ZSgxMjApIi8+CiAgPC9nPgo8L3N2Zz4K"></i> React
                </li>
                <li>
                    <i><img src="https://nodejs.org/static/favicon.ico"></i> Node
                </li>
                <li>
                    <i><img src="https://aws.amazon.com/favicon.ico"></i> AWS
                </li>
                <li>
                    <i><img src="https://cloud.google.com/favicon.ico"></i> Google Cloud
                </li>
            </ul>
        </section>

        <section class="experience">
            <h2>Experience</h2>

            <div class="res-item">
                <span class="title">Founder</span>
                <span class="organization"><a href="https://hexinsights.com">Hex Insights</a></span>
                <span class="dates">Feb 2018&mdash;Present</span>

                <ul>
                    <li>
                        Lead a small team of developers to build and deploy custom software that is tested, maintainable, and exceeds clients' expectations. Focusing, on systems to automate workflows by integrating new functionality and existing apps
                    </li>

                    <li>
                        Maintain client relationships and explore new business opportunities
                    </li>
                </ul>
            </div>

            <div class="res-item">
                <span class="title">CEO &amp; Co-Founder</span>
                <span class="organization">Honeyfire, Inc.</span>
                <span class="dates">Jan 2016&mdash;Feb 2018</span>

                <ul>
                    <li>
                        Created and marketed, with a team of two others, an event management platform for universities’ students and administrators called Flypost
                    </li>

                    <li>
                        Measured user engagement with Google Firebase and analyzed the generated data to guide decision making
                    </li>

                    <li>
                        Presented our product at the 2017 Educause Annual Conference in Philadelphia
                    </li>
                </ul>
            </div>

            <div class="res-item">
                <span class="title">App Development Mentor</span>
                <span class="organization">Foundry10</span>
                <span class="dates">Jun 2017&mdash;Nov 2017</span>

                <ul>
                    <li>
                        Taught small groups (2 to 4) of high school students how to build Android and Web applications
                    </li>

                    <li>
                        Facilitated learning of basic concepts of project planning, programming, and working as a team
                    </li>

                    <li>
                        Mentored students through difficult issues to reach solutions on their own
                    </li>
                </ul>
            </div>
        </section>
    </main>
</template>

<script>
export default {
    name: "Resume",
}
</script>

<style scoped>
h1 {
    font-size: 3rem;
    width: 100vw;
    margin-bottom: .5rem;
}

h2 {
    font-weight: 300;
    margin: .25rem auto .25rem;
}

ul {
    margin: 0;
}

a {
    text-decoration: none;
}

.title {
    display: block;
}

.organization, .dates {
    display: inline-block;
    margin-bottom: .3rem;
}

.dates {
    position: absolute;
    right: 2.5rem;
}

.res-item {
    margin-bottom: .75rem;
    padding: 0 .75rem;
}

.res-item .title {
    font-size: 1.2rem;
    font-weight: 500;
}

.res-item .dates {
    font-style: italic;
}

.col-container {
    display: flex;
    flex-flow: row;
}

.col-container section:first-of-type {
    width: 45%;
    box-sizing: border-box;
}

section.profiles {
    padding-left: 1.75rem;
    margin-top: 1.5rem;

    transition: var(--dark-transition);
}

@media print {
    section.profiles {
        margin-top: 0;
    }
}

section.profiles ul {
    list-style-type: none;
    padding: 0;
}

section.profiles i {
    display: inline-block;
    margin-right: .5rem
}

section.profiles i img {
    width: 1.5rem;
    transform: translate(0, .4rem);
}

section.education {
    margin-top: -.25rem;
}

section.education ul {
    padding: 0;
    display: flex;
    flex-flow: row;
    justify-content: space-around;
}

section.education ul li {
    text-align: center;
    display: block;
}

section.skills ul {
    list-style-type: none;
    padding: 0;
    display: flex;
    flex-flow: row;
    flex-wrap: wrap;
    justify-content: space-around;
    align-items: baseline;
}

section.skills ul li {
    text-align: center;
    position: relative;
    width: 20%;
    margin-bottom: .5rem;
}

section.skills i img {
    width: 1.5rem;
    display: block;
    margin: 0 auto .1rem;
}
</style>