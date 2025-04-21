<template>
  <ion-app>
    <ion-menu content-id="main-content">
      <ion-header>
        <ion-toolbar>
          <ion-title>{{ $t("menu.title") }}</ion-title>
        </ion-toolbar>
      </ion-header>
      <ion-content>
        <ion-item
          :button="true"
          class="cursor-pointer"
          @click="() => router.push('/')"
          >Home</ion-item
        >
        <ion-item :button="true" @click="() => router.push('/teamSettings')">{{
          $t("menu.teamSettings")
        }}</ion-item>
      </ion-content>
    </ion-menu>
    <ion-router-outlet id="main-content" />
  </ion-app>
</template>

<script setup lang="ts">
import {
  IonApp,
  IonContent,
  IonHeader,
  IonRouterOutlet,
  IonTitle,
  IonToolbar,
  IonMenu,
} from "@ionic/vue";
import { useRouter } from "vue-router";

const router = useRouter();
</script>
