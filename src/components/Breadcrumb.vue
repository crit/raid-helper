<template>
  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <router-link to="/xur">Xur</router-link>
    </li>

    <!--Home-->
    <li class="breadcrumb-item">
      <router-link to="/">Raids</router-link>
    </li>

    <!--Raid-->
    <li v-if="raid" class="breadcrumb-item">
      <router-link :to="`/${raid}`">{{ raidName }}</router-link>
    </li>

    <!--Encounter-->
    <li v-if="encounter" class="breadcrumb-item">
      <router-link :to="`/${raid}-${encounter}`">{{ encounterName }}</router-link>
    </li>
  </ol>
</template>

<script>
import {raids} from "@/data/raids";
import {encounters} from "@/data/encounters";

export default {
  name: "Breadcrumb",
  props: {
    raid: String,
    encounter: String
  },
  computed: {
    raidName() {
      const found = raids.find(raid => raid.id === this.raid)
      return found.name || 'unknown raid'
    },

    encounterName() {
      const found = encounters.find(encounter => encounter.id === this.encounter)
      return found.name || 'unknown encounter'
    }
  }
}
</script>

<style scoped>
</style>