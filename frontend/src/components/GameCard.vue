<template>
  <ion-grid :fixed="true">
    <ion-row
      ><ion-col>
        <ion-card>
          <ion-card-header>
            <ion-card-title class="flex flex-row justify-between">
              <span>{{ $props.league?.name }}</span>
              <span><ion-icon
          :icon="sportIcon"
        ></ion-icon></span>
            </ion-card-title>
          </ion-card-header>
          <ion-card-content>
            <div class="grid grid-cols-3 grid-rows-2">
              <div class="p-4 col-span-2">
                <div class="text-xs md:text-sm">{{ $t("GameCard.home") }}</div>
                <div class="md:text-lg">{{ $props.homeTeam?.name }}</div>
              </div>
              <div class="row-span-2 flex items-center justify-end">
                <div class="h-1/2 w-1/2 content-center text-center md:text-md">
                  <span class="timeBadge p-2 border-2 rounded-full">{{
                    startTime
                  }}</span>
                </div>
              </div>
              <div class="p-4 col-span-2">
                <div class="text-xs md:text-sm">{{ $t("GameCard.away") }}</div>
                <div class="md:text-lg">{{ $props.awayTeam?.name }}</div>
              </div>
            </div>
          </ion-card-content>
        </ion-card>
      </ion-col></ion-row
    >
  </ion-grid>
</template>

<script setup lang="ts">
import { IonGrid, IonRow, IonCol, IonCard, IonCardContent, IonIcon, IonCardTitle, IonCardHeader } from "@ionic/vue";
import { Game } from "../../gen/ts/kiioong/league_management/league_management";
import { format } from "date-fns";
import { inject, onMounted, ref } from "vue";
import { SERVICES } from "@/keys";

const props = defineProps<Game>();

const startDate = new Date(Number(props.startTimestamp) * 1000);

const startTime = format(startDate, "HH:mm");

const iconService = inject(SERVICES)?.iconService;

const sportIcon = ref<any>(null);

onMounted(async () => {
  sportIcon.value = await iconService?.loadIcon(props.league?.sport?.name + "Outline");
});



</script>

<style scoped>
ion-card {
  background-color: var(--ion-background-color);
}

ion-card-header {
  --tw-gradient-from: var(--ion-color-primary);
  --tw-gradient-to: var(--ion-color-secondary);
  --tw-gradient-stops:
    var(--tw-gradient-from), var(--tw-gradient-to, rgba(221, 214, 254, 0));
  background-image: linear-gradient(to right, var(--tw-gradient-stops));
}

ion-card-title {
  color: #fff;
}

.timeBadge {
  color: var(--ion-color-secondary);
}
</style>
