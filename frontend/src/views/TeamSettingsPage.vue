<template>
  <ion-page
    ><DefaultHeader></DefaultHeader>

    <ion-content>
      <ion-grid :fixed="true">
        <ion-row v-for="team in teams"
          ><ion-col>
            <ion-item>
              <ion-label>{{ team.name }} </ion-label>
              <ion-button
                slot="end"
                shape="round"
                color="danger"
                size="default"
                @click="() => removeTeamFromFavourites(team)"
              >
                <ion-icon slot="icon-only" :icon="trashOutline"></ion-icon
              ></ion-button>
            </ion-item> </ion-col
        ></ion-row>
      </ion-grid> </ion-content
  ></ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
  IonContent,
  IonItem,
  IonIcon,
  IonGrid,
  IonRow,
  IonCol,
  IonButton,
  IonLabel,
  modalController,
} from "@ionic/vue";
import { trashOutline } from "ionicons/icons";
import DefaultHeader from "@/components/DefaultHeader.vue";
import { inject, onMounted, ref, Ref } from "vue";
import { Team } from "../../gen/ts/kiioong/league_management/league_management";
import { SERVICES } from "@/keys";
import ConfirmModal from "@/components/ConfirmModal.vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const leagueManagementService = inject(SERVICES)?.leagueManagementService;

let teams: Ref<Team[]> = ref([]);

const removeTeamFromFavourites = async (team: Team) => {
  const modal = await modalController.create({
    component: ConfirmModal,
    componentProps: {
      title: t("TeamSettings.deleteTitle"),
      message: t("TeamSettings.deleteMessage"),
    },
  });

  await modal.present();

  const { data, role } = await modal.onWillDismiss();

  if (role === "confirm") {
    const success =
      await leagueManagementService?.removeTeamFromFavorites(team);

    if (!success) {
      console.error("team could not be removed");
    }

    await setFavouriteTeams();
  }
};

const setFavouriteTeams = async () => {
  teams.value = (await leagueManagementService?.getFavouriteTeams()) ?? [];
};

onMounted(async () => {
  await setFavouriteTeams();
});
</script>

<style scoped></style>
