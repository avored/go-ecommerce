import { createWebHistory, createRouter } from "vue-router"
import { RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/admin/login",
        name: "admin.login",
        component: () => import("./pages/auth/Login.vue"),
    },
    {
        path: "/tutorials/:id",
        name: "tutorial-details",
        component: () => import("./components/HelloWorld.vue"),
    },
    {
        path: "/",
        name: "add",
        component: () => import("./components/HelloWorld.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
