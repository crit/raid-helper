<template>
  <div class="list-group mt-5">
    <router-link
        v-for="encounter in encountersForRaid"
        :key="encounter.id"
        :to="`/${raid}-${encounter.id}`"
        :class="styler(encounter)">
      {{encounter.name}}
    </router-link>
  </div>
</template>

<script>
import {encounters} from "@/data/encounters";

export default {
  name: "EncounterChoices",
  props: {
    raid: String
  },
  computed: {
    encountersForRaid() {
      return encounters.filter(encounter => encounter.raid === this.raid)
    }
  },
  methods: {
    styler(encounter) {
      let style = 'list-group-item list-group-item-action'

      if (!encounter.active) {
        style += ' disabled'
      }

      return style
    }
  }
}
</script>

<style scoped>

</style>