import { createWebHistory, createRouter } from "vue-router"
import { RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/admin/login",
        name: "admin.login",
        component: () => import("./pages/auth/Login.vue"),
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: () => import("./pages/Dashboard.vue"),
    },
    {
        path: "/",
        name: "home",
        component: () => import("./components/HelloWorld.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
