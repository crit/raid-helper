<template>
  <Breadcrumb/>
  <RaidChoices/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import RaidChoices from "@/components/RaidChoices";

export default {
  name: 'RaidListView',
  components: {
    Breadcrumb,
    RaidChoices
  }
}
</script>
