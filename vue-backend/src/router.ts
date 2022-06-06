import { createWebHistory, createRouter } from "vue-router"
import { RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        alias: "/tutorials",
        name: "tutorials",
        component: () => import("./components/HelloWorld.vue"),
    },
    {
        path: "/tutorials/:id",
        name: "tutorial-details",
        component: () => import("./components/HelloWorld.vue"),
    },
    {
        path: "/add",
        name: "add",
        component: () => import("./components/HelloWorld.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
