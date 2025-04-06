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
            <ion-list>
              <ion-input
                v-model="username"
                :label="t('LoginPage.username')"
                label-placement="stacked"
                @keydown.enter="login"
              ></ion-input>
              <ion-input
                v-model="password"
                :label="t('LoginPage.password')"
                label-placement="stacked"
                type="password"
                @keydown.enter="login"
              >
                <ion-input-password-toggle
                  slot="end"
                ></ion-input-password-toggle
              ></ion-input>
              <ion-button buttonType="submit" @click="login">Login</ion-button>
            </ion-list>
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
  IonInput,
  IonInputPasswordToggle,
  IonList,
} from "@ionic/vue";
import { inject, ref } from "vue";
import { useRouter } from "vue-router";
import { SERVICES } from "@/keys";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
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

  const loginSuccessful = await authService?.login(
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
