import vuePlugin from "@vitejs/plugin-vue";
import laravel from "laravel-vite-plugin";
import { defineConfig } from "vite";

export default defineConfig({
    plugins: [
        vuePlugin(),
        laravel({
            input: [
                "resources/css/app.css",
                "resources/js/index.ts",
            ],
            refresh: true,
        }),
    ],
});
