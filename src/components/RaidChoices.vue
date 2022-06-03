<template>
  <div class="list-group mt-5">
    <router-link v-for="raid in raids" :key="raid.id" :to="raid.id" :class="styler(raid)">
      {{raid.name}}
    </router-link>
  </div>
</template>

<script>
import {raids} from "@/data/raids";

export default {
  name: "RaidChoices",
  computed: {
    raids() {
      return raids
    }
  },
  methods: {
    styler(raid) {
      let style = 'list-group-item list-group-item-action'

      if (!raid.active) {
        style += ' disabled'
      }

      return style
    }
  }
}
</script>

<style scoped>

</style>