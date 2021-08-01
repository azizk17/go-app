import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import axios from "axios";
import routes from "virtual:generated-pages";

import "./assets/style.css";
import { createAuth } from "@websanova/vue-auth";
import driverAuthBearer from "@websanova/vue-auth/dist/drivers/auth/bearer.esm.js";
import driverHttpAxios from "@websanova/vue-auth/dist/drivers/http/axios.1.x.esm.js";
import driverRouterVueRouter from "@websanova/vue-auth/dist/drivers/router/vue-router.2.x.esm.js";
import driverOAuth2Google from "@websanova/vue-auth/dist/drivers/oauth2/google.esm.js";
import driverOAuth2Facebook from "@websanova/vue-auth/dist/drivers/oauth2/facebook.esm.js";

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const auth = createAuth({
  plugins: {
    http: axios,
    router: router,
  },
  drivers: {
    http: driverHttpAxios,
    auth: driverAuthBearer,
    router: driverRouterVueRouter,
    oauth2: {
      google: driverOAuth2Google,
      facebook: driverOAuth2Facebook,
    },
  },
  options: {
    rolesKey: "type",
    notFoundRedirect: { name: "user-account" },
  },
});

createApp(App).use(router).use(auth).mount("#app");
