<template>
  <Breadcrumb raid="kf" encounter="kfc"/>
  <div class="text-center totem">
    <div v-for="(item, i) in items" :class="styler(i)">
      <img class="symbol-display" :src="item.path"/>
      <div class="text-center">{{item.name}}</div>
    </div>
  </div>
  <div class="text-center directions">
    <p v-for="direction in directions">{{direction}}</p>
  </div>
  <KfSymbolKeyboard v-on:symbol:selected="setSymbol" v-bind:only="allowed"/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import KfSymbolKeyboard from "@/components/kf/KfSymbolKeyboard";
import Storage from "@/data/storage";
import RingBuffer from "@/common/RingBuffer";

const version = "v0.1"

const descriptions = [
  "Below first swinging prison",
  "Hive ship jumping puzzle, near secret chest",
  "Annihilator Totem left room",
  "Warpriest right balcony",
  "Golgoroth maze final hole",
  "Golgoroth arena, lower left room",
  "Piston wall jumping puzzle, beside secret chest",
  "End of piston wall jumping puzzle, above final door",
  "In final room, above entrance",
]

// tuple of id to sort key.
const locations = {
  "hot": 0,
  "spider": 1,
  "ant": 2,
  "squid": 3,
  "sot": 4,
  "ufo": 5,
  "gerald": 6,
  "tesla": 7,
  "tongue": 8,
}

const emptySet = [
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "kf"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "kf"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "kf"}
]

export default {
  name: "EntranceView",
  components: {KfSymbolKeyboard, Breadcrumb},
  data: () => ({
    // localstorage table
    storage: Storage(`kf-kfc-${version}`),

    // entry buffer
    buffer: new RingBuffer([...emptySet]),

    // keyboard shows location keys and blank is prepended since it controls
    // resetting the entries.
    allowed: ['blank', ...Object.keys(locations)],
  }),
  created() {
    this.buffer = new RingBuffer(this.storage.read() || this.buffer.items())
  },
  computed: {
    directions() {
      return this.buffer.items()
        .map(item => item.id)
        .filter((v, i, a) => a.indexOf(v) === i)
        .map(id => locations[id])
        .sort()
        .map(id => descriptions[id])
    },

    items() {
      return this.buffer.items()
    }
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
  width: 80px;
  height: 80px;
}

.totem {
  display: grid;
  grid-template-columns: repeat(3, auto);
  max-width: 500px;
  margin: 30px auto 10px;
}

.directions {
  height: 175px;
  overflow: hidden;
}

.cursor img {
  border: 1px solid #444;
}
</style>