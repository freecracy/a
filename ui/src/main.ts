import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import ElementPlus from "element-plus";
import "element-plus/lib/theme-chalk/index.css";

import HomeApp from "./components/Home.vue";
import AboutApp from "./components/About.vue";

const Home = HomeApp;
const About = AboutApp;

const routes = [
  { path: "/", component: Home },
  { path: "/about", component: About },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).use(ElementPlus).mount("#app");
