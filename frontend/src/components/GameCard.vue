<template>
  <ion-grid :fixed="true">
    <div class="shadow-lg border-0 rounded gameCardWrapper">
      <ion-row class="ion-justify-content-center">
        <ion-col class="ion-text-center">
          {{ $props.homeTeam?.name }}
        </ion-col>
        <ion-col class="ion-text-center">{{ $props.awayTeam?.name }}</ion-col>
      </ion-row>
      <ion-row class="ion-justify-content-center">
        <ion-col :size="2" class="ion-text-center">{{ startTime }}</ion-col>
      </ion-row>
    </div>
  </ion-grid>
</template>

<script setup lang="ts">
import { IonGrid, IonRow, IonCol } from "@ionic/vue";
import { Game } from "../../gen/ts/kiioong/league_management/league_management";

const props = defineProps<Game>();

const startDate = new Date(Number(props.startTimestamp) * 1000);

const startTime = startDate.getHours() + ":" + startDate.getMinutes();
</script>

<style scoped>
.gameCardWrapper {
  background-color: var(--ion-color-primary);
  color: #e1e3e0;
}
</style>
