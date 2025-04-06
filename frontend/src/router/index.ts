import { createRouter, createWebHistory } from "@ionic/vue-router";
import { RouteRecordRaw } from "vue-router";
import { defineAsyncComponent, inject } from "vue";
import { Preferences } from "@capacitor/preferences";
import { SERVICES } from "@/keys";

const LoginPage = defineAsyncComponent(() => import("@/views/LoginPage.vue"));
const HomePage = () => import("@/views/HomePage.vue");
const TeamSettingsPage = () => import("@/views/TeamSettingsPage.vue");

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
    component: LoginPage,
  },
  {
    path: "/teamSettings",
    name: "TeamSettings",
    component: TeamSettingsPage,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

const authTokenResult = await Preferences.get({ key: "authToken" });

const isAuthenticated =
  authTokenResult.value !== null && authTokenResult.value !== "";

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
