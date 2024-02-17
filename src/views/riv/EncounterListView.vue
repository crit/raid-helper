<template>
  <Breadcrumb raid="riv"/>
  <EncounterChoices raid="riv"/>
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
