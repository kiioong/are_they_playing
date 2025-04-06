<template>
  <ion-content>
    <h1 class="pl-2 pb-2 border-b-1">{{ title }}</h1>
    <div class="text-center w-full absolute top-1/3">
      {{ message }}
    </div></ion-content
  >
  <ion-footer
    ><ion-grid>
      <ion-row class="ion-align-items-center">
        <ion-col class="text-center"
          ><ion-button @click="cancel" color="secondary">{{
            $t("actions.cancel")
          }}</ion-button></ion-col
        >
        <ion-col class="text-center"
          ><ion-button @click="confirm" color="danger"
            >Confirm</ion-button
          ></ion-col
        >
      </ion-row>
    </ion-grid></ion-footer
  >
</template>

<script setup lang="ts">
import { modalController } from "@ionic/vue";
import {
  IonContent,
  IonGrid,
  IonRow,
  IonCol,
  IonButton,
  IonFooter,
} from "@ionic/vue";

defineProps<{
  title: string;
  message: string;
}>();

const cancel = () => modalController.dismiss(null, "cancel");
const confirm = () => modalController.dismiss(null, "confirm");
</script>

<style scoped></style>
