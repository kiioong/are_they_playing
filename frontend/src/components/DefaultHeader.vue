<template>
  <ion-header :translucent="true">
    <ion-toolbar>
      <ion-buttons slot="start">
        <ion-menu-button></ion-menu-button>
      </ion-buttons>
      <ion-title>Are they Playing?</ion-title>
      <ion-buttons slot="end">
        <ion-button @click="authService?.logout()">Logout</ion-button>
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
import { useRouter } from "vue-router";
import { inject } from "vue";
import { SERVICES } from "@/keys";

const authService = inject(SERVICES)?.authService;
</script>

<style scoped></style>
