<template>
  <Breadcrumb raid="vod" encounter="exi"/>
  <p class="text-center">First two rooms skipped since they are easy.</p>
  <div class="text-center totem">
    <div v-for="(item, i) in items" :class="styler(i)">
      <img :src="item.path" :alt="item.id" class="symbol-display">
      <p class="text-center">{{item.name}}</p>
    </div>
  </div>
  <VodSymbolKeyboard v-on:symbol:selected="setSymbol" v-bind:exclude="excluded"/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import VodSymbolKeyboard from "@/components/vod/VodSymbolKeyboard";
import Storage from "@/data/storage";
import RingBuffer from "@/common/RingBuffer";

const version = "v0.1"

const emptySet = [
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
]

export default {
  name: "ExhibitionView",
  components: {VodSymbolKeyboard, Breadcrumb},
  data: () => ({
    storage: Storage(`vod-exi-${version}`),
    buffer: new RingBuffer([...emptySet]),
    excluded: []
  }),
  computed: {
    items() {
      return this.buffer.items()
    }
  },
  created() {
    this.buffer = new RingBuffer(this.storage.read() || this.buffer.items())
  },
  methods: {
    setSymbol(symbol) {
      if (symbol.id === 'blank') {
        this.buffer = new RingBuffer([...emptySet])
      } else {
        this.buffer.set(symbol)
      }

      this.storage.write(this.buffer.items())
    },

    styler(index = 0) {
      const styles = []

      if (this.buffer.cursor === index) styles.push('cursor')

      return styles.join(' ')
    }
  }
}
</script>

<style scoped>
.symbol-display {
  width: 50px;
  height: 50px;
}

.totem {
  display: grid;
  grid-template-columns: repeat(2, auto);
  max-width: 500px;
  margin: 30px auto 10px;
}

.cursor img {
  border: 1px solid #444;
}
</style>