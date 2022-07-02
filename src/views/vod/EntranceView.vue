<template>
  <Breadcrumb raid="vod" encounter="ent"/>
  <div class="text-center totem">
    <div v-for="(item, i) in items" :class="styler(i)">
      <img class="symbol-display" :src="item.path"/>
      <div class="text-center">{{item.name}}</div>
    </div>
  </div>
  <div class="text-center directions">
    <p v-for="direction in directions">{{direction}}</p>
  </div>
  <VodSymbolKeyboard v-on:symbol:selected="setSymbol" v-bind:only="allowed"/>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import VodSymbolKeyboard from "@/components/vod/VodSymbolKeyboard";
import Storage from "@/data/storage";
import RingBuffer from "@/common/RingBuffer";

const version = "v0.1"

const descriptions = [
  "In the first room after you drop down when you first enter the pyramid",
  "In the room just after the room with the frozen Scorn, before the first encounter",
  "In the first encounter room",
  "In the room with the purple wall directly before the second encounter",
  "In the final stand area of the second encounter",
  "In the jumping puzzle after the second encounter, in the room where the first Darkness Crux used for the jumping puzzle is",
  "Further along in the jumping puzzle, after the second encounter",
  "To the left immediately after the third encounter",
  "Before the final boss, where you have to climb up the side walls",
]

// tuple of id to sort key.
const locations = {
  "pyramid": 0,
  "give": 1,
  "dark": 2,
  "traveler": 3,
  "worship": 4,
  "light": 5,
  "stop": 6,
  "guardian": 7,
  "kill": 8,
}

const emptySet = [
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"},
  {id: "blank", name: "Blank", path: "assets/vod/blank.png", raid: "vod"}
]

export default {
  name: "EntranceView",
  components: {VodSymbolKeyboard, Breadcrumb},
  data: () => ({
    // localstorage table
    storage: Storage(`vod-ent-${version}`),

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