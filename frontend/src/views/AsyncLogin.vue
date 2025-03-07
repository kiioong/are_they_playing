<script setup lang="ts">

import Login from "@/views/Login.vue";
</script>

<template>
  <Suspense>
    <Login></Login>
    <template #fallback>
      Loading
    </template>
  </Suspense>
</template>

<style scoped>

</style>