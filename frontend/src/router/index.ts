import { createRouter, createWebHistory } from "@ionic/vue-router";
import { RouteRecordRaw } from "vue-router";
import { defineAsyncComponent, inject } from "vue";
import { Preferences } from "@capacitor/preferences";
import { SERVICES } from "@/keys";

const Login = defineAsyncComponent(() => import("@/views/Login.vue"));
const HomePage = () => import("@/views/HomePage.vue");

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/home",
  },
  {
    path: "/home",
    name: "Home",
    component: HomePage,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

const authTokenResult = await Preferences.get({ key: "authToken" });

const isAuthenticated =
  authTokenResult.value !== null && authTokenResult.value !== "";

console.log(authTokenResult, isAuthenticated);

router.beforeEach((to) => {
  const authService = inject(SERVICES)?.authService;
  authService?.setToken(authTokenResult.value ?? "");

  if (
    // make sure the user is authenticated
    !isAuthenticated &&
    to.name !== "Login"
  ) {
    // redirect the user to the login page
    return { name: "Login" };
  }
});

export default router;
