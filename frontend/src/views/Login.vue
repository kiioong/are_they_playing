<template>
  <IonPage>
    <ion-header :translucent="true">
      <ion-toolbar>
        <ion-title>Login</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content>
      <form name="login-form" >
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
    </ion-content>

  </IonPage>

</template>

<script setup lang="ts">
import {IonContent, IonHeader, IonPage, IonTitle, IonToolbar, IonButton} from '@ionic/vue';
import {AuthenticationClient} from '@/../gen/ts/kiioong/authentication/authentication_service.client'
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {ref} from "vue";
import {Preferences} from "@capacitor/preferences";
import {useRouter} from "vue-router";

const router = useRouter();

const username = ref('');
const password = ref('');


const login = async () => {
  if (!username.value || !password.value) {
    //validation
    console.error('Please enter a username and password');
    return;
  }

  const transport = new GrpcWebFetchTransport({
    baseUrl: "http://localhost:10000",
  });

  const ac = new AuthenticationClient(transport);
  const result = await ac.login({username: username.value, password: password.value})

  if (result.response.jwtToken === '') {
    // validation
    console.error('Username not found or wrong password');
    return;
  }

  await Preferences.set({key: "authToken", value: result.response.jwtToken});

  await router.push('/home');
}

</script>