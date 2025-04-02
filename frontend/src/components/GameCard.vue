<template>
  {{ $props.homeTeam?.name }} - {{ $props.awayTeam?.name }} -
  {{ startTime }}
</template>

<script setup lang="ts">
import { Game } from "../../gen/ts/kiioong/league_management/league_management";

const props = defineProps<Game>();

const startDate = new Date(Number(props.startTimestamp) * 1000);

const startTime = startDate.getHours() + ":" + startDate.getMinutes();

console.log(props);
</script>

<style scoped></style>
