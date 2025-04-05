<template>
  <ion-grid class="h-7 p-1">
    <ion-row>
      <ion-col class="text-start">
        <div @click="decreaseDay" class="h-8 w-8 absolute left-0 -top-0.5">
          <ion-icon class="m-2" :icon="caretBackOutline"></ion-icon></div
      ></ion-col>
      <ion-col class="text-center">{{
        isSameDay(pickedDay, new Date())
          ? $t("misc.today")
          : format(pickedDay, "dd.MM.y")
      }}</ion-col>
      <ion-col class="text-end">
        <div @click="increaseDay" class="h-8 w-8 absolute right-0 -top-0.5">
          <ion-icon class="m-2" :icon="caretForwardOutline"></ion-icon>
        </div>
      </ion-col>
    </ion-row>
  </ion-grid>
</template>

<script setup lang="ts">
import { IonIcon, IonGrid, IonRow, IonCol } from "@ionic/vue";
import { caretBackOutline, caretForwardOutline } from "ionicons/icons";
import { format, addDays, isSameDay } from "date-fns";
import { ModelRef } from "vue";

const pickedDay: ModelRef<Date> = defineModel();

const increaseDay = () => {
  pickedDay.value = addDays(pickedDay.value, 1);
};

const decreaseDay = () => {
  pickedDay.value = addDays(pickedDay.value, -1);
};
</script>

<style scoped>
ion-grid {
  padding: 0;
  background-color: var(--ion-color-secondary);
  color: #e1e3e0;
}
</style>
