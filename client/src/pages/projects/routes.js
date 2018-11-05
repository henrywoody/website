import Drobe from './Drobe.vue';
import Website from './Website.vue';
import Holland from './Holland.vue';
import Boids from './Boids.vue';
import ODCA from './ODCA.vue';
import TDCA from './TDCA.vue';

export default [
    {path: "/projects/drobe", component: Drobe},
    {path: "/projects/website", component: Website},
    {path: "/projects/holland", component: Holland},
    {path: "/projects/boids", component: Boids},
    {path: "/projects/1d-cellular-automata", component: ODCA},
    {path: "/projects/2d-cellular-automata", component: TDCA},
]