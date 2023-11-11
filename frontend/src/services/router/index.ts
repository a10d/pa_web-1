import {createRouter, createWebHistory} from "vue-router";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "home",
            component: () => import("../../views/Home.vue"),
        },
        {
            path: "/admin",
            name: "admin",
            component: () => import("../../views/Admin.vue"),
        },
    ],
});
