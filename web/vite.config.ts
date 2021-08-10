import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import Pages from "vite-plugin-pages";
import ViteComponents from "vite-plugin-components";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ViteComponents(),
    Pages({
      pagesDir: "src/pages",
      // pagesDir: [
      //   { dir: "src/pages", baseRoute: "" },
      //   { dir: "src/features/**/pages", baseRoute: "features" },
      //   { dir: "src/admin/pages", baseRoute: "admin" },
      // ],
    }),
  ],
  server: {
    proxy: {
      '/api': {
        target: 'http://app:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
    },
    hmr: {
      
    },
    watch: {
      usePolling: true,
    },
  },
});
