<template>
  <IonPage>
    <ion-header>
      <ion-toolbar>
        <ion-title>Login</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content>
      <ion-grid :fixed="true">
        <ion-row class="ion-justify-content-center">
          <ion-col size="3">
            <form name="login-form">
              <div class="mb-3">
                <label for="username">Username: </label>
                <input v-model="username" id="username" type="text" />
              </div>
              <div class="mb-3">
                <label for="password">Password: </label>
                <input v-model="password" id="password" type="password" />
              </div>
              <ion-button @click="login">Login</ion-button>
            </form>
          </ion-col>
        </ion-row>
      </ion-grid>
    </ion-content>
  </IonPage>
</template>

<script setup lang="ts">
import {
  IonContent,
  IonHeader,
  IonPage,
  IonTitle,
  IonToolbar,
  IonButton,
  IonGrid,
  IonCol,
  IonRow,
} from "@ionic/vue";
import { inject, ref } from "vue";
import { useRouter } from "vue-router";
import { SERVICES } from "@/keys";

const router = useRouter();

const username = ref("");
const password = ref("");

const authService = inject(SERVICES)?.authService;

const login = async () => {
  if (!username.value || !password.value) {
    //validation
    console.error("Please enter a username and password");
    return;
  }

  const loginSuccessful = await authService.login(
    username.value,
    password.value,
  );

  console.log(loginSuccessful);

  if (!loginSuccessful) {
    return;
  }

  await router.push("/home");
};
</script>
