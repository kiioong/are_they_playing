<template>
  <ion-header>
    <ion-toolbar>
      <ion-buttons slot="start">
        <ion-button @click="$emit('closeModal')">{{
          $t("actions.cancel")
        }}</ion-button>
      </ion-buttons>
      <ion-title>{{ $t("TeamSearch.caption") }}</ion-title>
    </ion-toolbar>
  </ion-header>
  <ion-content class="ion-padding">
    <ion-item>
      <ion-input
        :label="$t('TeamSearch.search')"
        label-placement="stacked"
        ref="input"
        type="text"
        :placeholder="$t('TeamSearch.placeholder')"
      ></ion-input>
    </ion-item>

    <ion-item
      v-if="leagues.length === 0"
      v-for="sport in sports"
      @click="selectSport(sport)"
    >
      <ion-icon
        class="ion-margin-end"
        :icon="getIcon(sport.name + 'Outline')"
      ></ion-icon>
      {{ $t("sports." + sport.name) }}
    </ion-item>
    <ion-item
      v-if="teams.length === 0"
      v-for="league in leagues"
      @click="selectLeague(league)"
    >
      {{ league.name }}
    </ion-item>
    <ion-item v-for="team in teams" @click="selectTeam(team)">
      {{ team.name }}
    </ion-item>
  </ion-content>
</template>

<script setup lang="ts">
import {
  IonContent,
  IonHeader,
  IonTitle,
  IonToolbar,
  IonItem,
  IonInput,
  IonButtons,
  IonButton,
  IonIcon,
} from "@ionic/vue";
import {
  League,
  Sport,
  Team,
} from "../../gen/ts/kiioong/league_management/league_management";
import { inject, onMounted, reactive, Ref, ref } from "vue";
import { SERVICES } from "@/keys";

const emit = defineEmits(["closeModal"]);

const leagueManagementService = inject(SERVICES)?.leagueManagementService;

let sports: Ref<Sport[]> = ref([]);
let leagues: Ref<League[]> = ref([]);
let teams: Ref<Team[]> = ref([]);

onMounted(async () => {
  sports.value = (await leagueManagementService?.getSports()) ?? [];
});

const selectSport = async (sport: Sport) => {
  leagues.value = (await leagueManagementService?.getLeagues(sport)) ?? [];
};

const selectLeague = async (league: League) => {
  teams.value = (await leagueManagementService?.getTeams(league)) ?? [];
};

const selectTeam = async (team: Team) => {
  const successful =
    (await leagueManagementService?.addTeamToFavorites(team)) ?? [];

  if (successful) {
    emit("closeModal");
  }
};

// Store the loaded icons in a cache
const iconCache = reactive<Record<string, any>>({});
// Store promise map for pending icon loads to prevent duplicate requests
const loadingIcons = reactive<Record<string, Promise<any>>>({});

// Function to load an icon and cache it
const loadIcon = async (iconName: string): Promise<any> => {
  if (!iconName) return null;

  // If this icon is already being loaded, wait for that promise
  if (await loadingIcons[iconName]) {
    return loadingIcons[iconName];
  }

  // If the icon is in cache, return it immediately
  if (iconCache[iconName]) {
    return iconCache[iconName];
  }

  // Otherwise, load the icon and cache it
  const iconPromise = import(`ionicons/icons/`)
    .then((module) => {
      const icon = module[iconName];
      iconCache[iconName] = icon;
      delete loadingIcons[iconName];
      return icon;
    })
    .catch(async (error) => {
      console.error(`Failed to load icon: ${iconName}`, error);
      // Load a default fallback icon
      return import("ionicons/icons").then((module) => {
        const icon = module.help;
        iconCache[iconName] = icon;
        delete loadingIcons[iconName];
        return icon;
      });
    });

  // Store the promise in the loading map
  loadingIcons[iconName] = iconPromise;
  return iconPromise;
};

// Function to get an icon (will return immediately if cached, or null during loading)
const getIcon = (iconName: string): any => {
  // Return cached icon if available
  if (iconCache[iconName]) {
    return iconCache[iconName];
  }

  // Start loading if not already loading
  if (!loadingIcons[iconName]) {
    loadIcon(iconName);
  }

  // Return a default/fallback icon while loading
  return iconCache["default"];
};
</script>

<style scoped></style>
