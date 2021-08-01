import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import Pages from "vite-plugin-pages";
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Pages({
      pagesDir: "src/pages",
      // pagesDir: [
      //   { dir: "src/pages", baseRoute: "" },
      //   { dir: "src/features/**/pages", baseRoute: "features" },
      //   { dir: "src/admin/pages", baseRoute: "admin" },
      // ],
    }),
  ],
});
