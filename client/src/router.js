import Vue from 'vue';
import Router from 'vue-router';
import Index from './pages/Index.vue';
import Projects from './pages/Projects.vue';
import Links from './pages/Links.vue';
import Resume from './pages/Resume.vue';

import projectRoutes from './pages/projects/routes';

Vue.use(Router);

const baseRoutes = [
    {path: "/", component: Index},
    {path: "/projects", component: Projects},
    {path: "/links", component: Links},
    {path: "/resume", component: Resume},
]

export default new Router({
    mode: "history",
    routes: [
        ...baseRoutes,
        ...projectRoutes,
    ],
    scrollBehavior() {
        return {x: 0, y: 0}
    }
});
