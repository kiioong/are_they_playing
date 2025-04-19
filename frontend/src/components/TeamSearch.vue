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
    <div v-if="leagues.length === 0">
      <ion-item
        v-for="sport in sports"
        :key="sport.id"
        @click="selectSport(sport)"
      >
        <ion-icon class="ion-margin-end" :icon="sport.icon"></ion-icon>
        {{ $t("sports." + sport.name) }}
      </ion-item>
    </div>

    <div v-if="teams.length === 0">
      <ion-item
        v-for="league in leagues"
        :key="league.id"
        @click="selectLeague(league)"
      >
        {{ league.name }}
      </ion-item>
    </div>

    <ion-item v-for="team in teams" :key="team.id" @click="selectTeam(team)">
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
import { inject, onMounted, Ref, ref } from "vue";
import { SERVICES } from "@/keys";

const emit = defineEmits(["closeModal"]);

const leagueManagementService = inject(SERVICES)?.leagueManagementService;

interface SportWithIcon extends Sport {
  icon?: any;
}

let sports: Ref<SportWithIcon[]> = ref([]);
let leagues: Ref<League[]> = ref([]);
let teams: Ref<Team[]> = ref([]);

onMounted(async () => {
  sports.value = (await leagueManagementService?.getSports()) ?? [];
  sports.value.map(async (sport) => {
    sport.icon = await iconService?.loadIcon(sport.name + "Outline");
    return sport;
  });
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

const iconService = inject(SERVICES)?.iconService;
</script>

<style scoped></style>
