<template>
  <ion-page>
    <DefaultHeader></DefaultHeader>
    <ion-content id="main-content" :fullscreen="true">
      <DayToggleBar v-model="pickedDay"></DayToggleBar>
      <ion-grid :fixed="true">
        <ion-row>
          <ion-col>
            <template v-if="games.length > 0">
              <div class="grid grid-cols-1 lg:grid-cols-2">
                <GameCard
                  v-for="game in games"
                  :key="game.homeTeam?.id"
                  :start-timestamp="game.startTimestamp"
                  :home-team="game.homeTeam"
                  :away-team="game.awayTeam"
                  :league="game.league"
                ></GameCard>
              </div>
            </template>
            <template v-else>
              <ion-card
                ><ion-card-content class="text-center">{{
                  $t("HomePage.noGamesMessage")
                }}</ion-card-content></ion-card
              >
            </template>
          </ion-col>
        </ion-row>
      </ion-grid>

      <ion-fab slot="fixed">
        <ion-fab-button id="open-add-team-modal">
          <ion-icon :icon="addOutline"></ion-icon>
        </ion-fab-button>
      </ion-fab>
      <ion-modal ref="modal" trigger="open-add-team-modal">
        <TeamSearch @close-modal="cancel" @team-selected="teamSelected"></TeamSearch>
      </ion-modal>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonContent,
  IonPage,
  IonIcon,
  IonModal,
  IonFab,
  IonFabButton,
  IonGrid,
  IonRow,
  IonCol,
  IonCard,
  IonCardContent,
} from "@ionic/vue";
import { addOutline } from "ionicons/icons";
import { defineAsyncComponent, inject, onMounted, Ref, ref, watch } from "vue";
import { Game } from "../../gen/ts/kiioong/league_management/league_management";
import { SERVICES } from "@/keys";
import DayToggleBar from "@/components/DayToggleBar.vue";
import DefaultHeader from "@/components/DefaultHeader.vue";

const TeamSearch = defineAsyncComponent(
  () => import("@/components/TeamSearch.vue"),
);

const GameCard = defineAsyncComponent(
  () => import("@/components/GameCard.vue"),
);

const leagueManagementService = inject(SERVICES)?.leagueManagementService;

const modal = ref();
let games: Ref<Game[]> = ref([]);
const pickedDay: Ref<Date> = ref(new Date());

const cancel = () => modal.value.$el.dismiss(null, "cancel");

onMounted(async () => {
  await getGames(pickedDay.value);
});

watch(pickedDay, async (newDay) => {
  await getGames(newDay);
});

const getGames = async (pickedDay: Date) => {
  games.value = (await leagueManagementService?.getGames(pickedDay)) ?? [];
};

const teamSelected = async () => {
  modal.value.$el.dismiss(null, "teamSelected");
  await getGames(pickedDay.value);
};
</script>

<style scoped>
ion-fab[slot="fixed"] {
  top: 90%;
  right: 20px;
}
</style>
