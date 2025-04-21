<template>
  <IonPage>
    <ion-header>
      <ion-toolbar>
        <ion-title>Login</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content>
      <div class="flex h-[70vh] items-center justify-center">
        <ion-grid :fixed="true">
          <ion-row class="ion-justify-content-center">
            <ion-col class="!max-w-[330px]" size="10" size-md="6" size-lg="3">
              <ion-card>
                <ion-card-content>
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

                    <ion-button
                      class="w-full rounded-full pr-[2px]"
                      size="small"
                      color="primary"
                      @click="login"
                      >Login</ion-button
                    >
                  </ion-list>
                </ion-card-content>
              </ion-card>
            </ion-col>
          </ion-row>
        </ion-grid>
      </div>
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
  IonCard,
  IonCardContent,
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

  await router
    .push({ path: "/home" })
    .then(() => {
      console.log("Navigation successful");
    })
    .catch((error) => {
      console.error("Error navigating to home:", error);
    });
  // router.go(0);

  console.log("Login successful");
};
</script>
