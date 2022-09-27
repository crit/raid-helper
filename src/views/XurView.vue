<template>
  <Breadcrumb/>
  <div class="card mt-5">
    <div class="card-body">
      <h4 class="card-title">Xur Location</h4>
      <p class="card-text">{{locationName}}</p>
    </div>
  </div>

  <div class="card mt-5">
    <div class="card-body">
      <h4 class="card-title">Xur Inventory</h4>
      <p class="card-text"><em>Not Yet Implemented</em></p>
    </div>
  </div>
</template>

<script>
import {xur, Location} from "@/data/xur";
import Breadcrumb from "@/components/Breadcrumb";

export default {
  name: "XurView",
  components: {Breadcrumb},
  data: () => ({
    location: new Location()
  }),
  async created() {
    this.location = await xur.location()
  },
  computed: {
    locationName() {
      return this.location.expired()
          ? `Xur has left: ${this.location.locationName}`
          : `Xur is at: ${this.location.locationName}`
    },
  }
}
</script>

<style scoped>

</style>
