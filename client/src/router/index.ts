import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Login from "../pages/Login.vue";
import Register from "../pages/Register.vue";
import { useRedux } from "../lib/hooks/useRedux";

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
    beforeEnter: (to, from, next) => {
      const isLoggedIn = useRedux((state) => state.auth.isLoggedIn).value;
      if (!isLoggedIn) {
        next({ name: "login" });
      } else {
        next();
      }
    },
  },
  {
    path: "/accounts/login",
    name: "login",
    component: Login,
  },
  {
    path: "/accounts/register",
    name: "register",
    component: Register,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
