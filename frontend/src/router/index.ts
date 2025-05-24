import { createRouter, createWebHistory } from "@ionic/vue-router";
import { RouteRecordRaw } from "vue-router";
import { inject } from "vue";
import { Preferences } from "@capacitor/preferences";
import { SERVICES } from "@/keys";
import { AuthService } from "@/services/auth.service";

const LoginPage = () => import("@/views/LoginPage.vue");
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

const isAuthenticated = async () => {
  const authTokenResult = await Preferences.get({ key: "authToken" });
  if (!authTokenResult.value) return false;

  const authService = new AuthService();

  if (!authService) return false;

  return await authService.validateToken(authTokenResult.value);
};

router.beforeEach(async (to, from, next) => {
  const authService = inject(SERVICES)?.authService;
  authService?.setToken(
    (await Preferences.get({ key: "authToken" })).value ?? "",
  );

  if (
    // make sure the user is authenticated
    !(await isAuthenticated()) &&
    to.name !== "Login"
  ) {
    // redirect the user to the login page
    next("/login");
    return;
  }

  next();
});

export default router;
