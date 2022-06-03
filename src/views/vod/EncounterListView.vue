<template>
  <Breadcrumb raid="vod"/>
  <EncounterChoices raid="vod"/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import EncounterChoices from "@/components/EncounterChoices";

export default {
  name: "EncounterListView",
  components: {EncounterChoices, Breadcrumb},
}
</script>

<style scoped>

</style>