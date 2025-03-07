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
          <input id="username" type="text" />
        </div>
        <div class="mb-3">
          <label for="password">Password: </label>
          <input id="password" type="password" />
        </div>
        <ion-button @click="login">Login</ion-button>
      </form>
    </ion-content>

  </IonPage>

</template>

<script setup lang="ts">
import {IonContent, IonHeader, IonPage, IonTitle, IonToolbar} from '@ionic/vue';
import {AuthentificationClient} from '@/../gen/ts/kiioong/authentication_service/authentication_service.client'
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";



const login = async () => {
  const transport = new GrpcWebFetchTransport({
    baseUrl: "http://localhost:10000",
  });

  const ac = new AuthentificationClient(transport);
  const result = await ac.login({username: 'Admin', password: '12345', sessionId: -1})
  console.log(result.response);
}


</script>