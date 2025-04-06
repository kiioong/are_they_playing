<template>
  <ion-app>
    <ion-menu content-id="main-content">
      <ion-header>
        <ion-toolbar>
          <ion-title>{{ $t("menu.title") }}</ion-title>
        </ion-toolbar>
      </ion-header>
      <ion-content class="ion-padding">
        <div class="cursor-pointer mb-2" @click="() => router.push('/')">
          Home
        </div>
        <div
          class="cursor-pointer mb-2"
          @click="() => router.push('/teamSettings')"
        >
          {{ $t("menu.teamSettings") }}
        </div>
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
