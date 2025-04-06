<template>
  <ion-header :translucent="true">
    <ion-toolbar>
      <ion-buttons slot="start">
        <ion-menu-button></ion-menu-button>
      </ion-buttons>
      <ion-title>Are they Playing?</ion-title>
      <ion-buttons slot="end">
        <ion-button @click="logout()">Logout</ion-button>
      </ion-buttons>
    </ion-toolbar>
  </ion-header>
</template>

<script setup lang="ts">
import {
  IonButton,
  IonButtons,
  IonHeader,
  IonTitle,
  IonToolbar,
  IonMenuButton,
} from "@ionic/vue";
import { Preferences } from "@capacitor/preferences";
import { useRouter } from "vue-router";

const router = useRouter();

const logout = async () => {
  await Preferences.remove({ key: "authToken" });
  await router.push("/login");
};
</script>

<style scoped></style>
