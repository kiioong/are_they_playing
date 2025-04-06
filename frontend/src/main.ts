import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createI18n } from "vue-i18n";
import ServicesPlugin from "@/plugins/services";

import { IonicVue } from "@ionic/vue";

/* Core CSS required for Ionic components to work properly */
import "@ionic/vue/css/core.css";

/* Basic CSS for apps built with Ionic */
import "@ionic/vue/css/normalize.css";
import "@ionic/vue/css/structure.css";
import "@ionic/vue/css/typography.css";

/* Optional CSS utils that can be commented out */
import "@ionic/vue/css/padding.css";
import "@ionic/vue/css/float-elements.css";
import "@ionic/vue/css/text-alignment.css";
import "@ionic/vue/css/text-transformation.css";
import "@ionic/vue/css/flex-utils.css";
import "@ionic/vue/css/display.css";

/**
 * Ionic Dark Mode
 * -----------------------------------------------------
 * For more info, please see:
 * https://ionicframework.com/docs/theming/dark-mode
 */

/* @import '@ionic/vue/css/palettes/dark.always.css'; */
/* @import '@ionic/vue/css/palettes/dark.class.css'; */
// import "@ionic/vue/css/palettes/dark.system.css";
/* Theme variables */
import "./theme/variables.css";

const i18n = createI18n({
  legacy: false,
  locale: "de",
  fallbackLocale: "en",
  messages: {
    en: {
      actions: {
        cancel: "CANCEL",
      },
      menu: {
        title: "Menu",
        teamSettings: "Team Settings",
      },
      sports: {
        football: "Football",
      },
      TeamSearch: {
        caption: "Choose a team to add",
        search: "Search",
        placeholder: "sport/league/team",
      },
      TeamSettings: {
        deleteTitle: "Remove Team",
        deleteMessage: "Are you sure you want to remove this team?",
      },
      misc: {
        today: "Today",
      },
    },
    de: {
      actions: {
        cancel: "ABBRECHEN",
      },
      menu: {
        title: "Menü",
        teamSettings: "Team Einstellungen",
      },
      sports: {
        football: "Fussball",
      },
      TeamSearch: {
        caption: "Wähle ein Team zum Hinzufügen",
        search: "Suche",
        placeholder: "Sport/Liga/Team",
      },
      TeamSettings: {
        deleteTitle: "Team Entfernen",
        deleteMessage: "Möchtest du dieses Team wirklich entfernen?",
      },
      misc: {
        today: "Heute",
      },
    },
  },
});

const app = createApp(App)
  .use(IonicVue)
  .use(router)
  .use(i18n)
  .use(ServicesPlugin);

router.isReady().then(() => {
  app.mount("#app");
});
