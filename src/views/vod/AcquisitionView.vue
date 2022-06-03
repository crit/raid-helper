<template>
  <Breadcrumb raid="vod" encounter="acq"/>
  <div class="text-center totem">
    <div v-for="item in buffer.items()" :key="item.id">
      <img class="symbol-display" :src="item.path"/>
      <div class="text-center">{{item.name}}</div>
    </div>
  </div>
  <VodSymbolKeyboard v-on:symbol:selected="setSymbol"/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import Storage from "@/data/storage";
import RingBuffer from "@/common/RingBuffer";
import VodSymbolKeyboard from "@/components/vod/VodSymbolKeyboard";

export default {
  name: "AcquisitionView",
  components: {Breadcrumb, VodSymbolKeyboard},
  data: () => ({
    storage: Storage("vod-acq"),
    buffer: new RingBuffer([
      {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
      {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
      {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"}
    ]),
  }),
  created() {
    this.buffer = new RingBuffer(this.storage.read() || this.buffer.items())
  },
  methods: {
    setSymbol(symbol) {
      this.buffer.set(symbol)
      this.storage.write(this.buffer.items())
    }
  }
}
</script>

<style scoped>
.symbol-display {
  width: 80px;
  height: 80px;
}

.totem {
  display: grid;
  grid-template-columns: repeat(3, auto);
  max-width: 500px;
  margin: 30px auto 10px;
}
</style>