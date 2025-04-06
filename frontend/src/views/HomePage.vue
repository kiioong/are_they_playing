<template>
  <ion-page>
    <DefaultHeader></DefaultHeader>
    <ion-content id="main-content" :fullscreen="true">
      <DayToggleBar v-model="pickedDay"></DayToggleBar>
      <GameCard
        v-for="game in games"
        :key="game.homeTeam?.id"
        :start-timestamp="game.startTimestamp"
        :home-team="game.homeTeam"
        :away-team="game.awayTeam"
      ></GameCard>
      <ion-fab slot="fixed">
        <ion-fab-button id="open-add-team-modal">
          <ion-icon :icon="addOutline"></ion-icon>
        </ion-fab-button>
      </ion-fab>
      <ion-modal ref="modal" trigger="open-add-team-modal">
        <TeamSearch @close-modal="cancel"></TeamSearch>
      </ion-modal>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonContent,
  IonHeader,
  IonPage,
  IonTitle,
  IonToolbar,
  IonIcon,
  IonModal,
  IonFab,
  IonFabButton,
  IonButton,
  IonButtons,
} from "@ionic/vue";
import { Preferences } from "@capacitor/preferences";
import { useRouter } from "vue-router";
import { addOutline } from "ionicons/icons";
import { defineAsyncComponent, inject, onMounted, Ref, ref, watch } from "vue";
import { Game } from "../../gen/ts/kiioong/league_management/league_management";
import { SERVICES } from "@/keys";
import GameCard from "@/components/GameCard.vue";
import DayToggleBar from "@/components/DayToggleBar.vue";
import DefaultHeader from "@/components/DefaultHeader.vue";

const TeamSearch = defineAsyncComponent(
  () => import("@/components/TeamSearch.vue"),
);

const leagueManagementService = inject(SERVICES)?.leagueManagementService;

const modal = ref();
let games: Ref<Game[]> = ref([]);
const pickedDay: Ref<Date> = ref(new Date());

const cancel = () => modal.value.$el.dismiss(null, "cancel");

onMounted(async () => {
  games.value =
    (await leagueManagementService?.getGames(pickedDay.value)) ?? [];
});

watch(pickedDay, async (newDay) => {
  games.value = (await leagueManagementService?.getGames(newDay)) ?? [];
});
</script>

<style scoped>
ion-fab[slot="fixed"] {
  top: 90%;
  right: 20px;
}
</style>
