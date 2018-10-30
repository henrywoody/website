import Vue from 'vue';
import Router from 'vue-router';
import Index from './pages/Index.vue';
import Projects from './pages/Projects.vue';

Vue.use(Router);

export default new Router({
    mode: "history",
    routes: [
        {path: "/", component: Index},
        {path: "/projects", component: Projects}
    ]
});
