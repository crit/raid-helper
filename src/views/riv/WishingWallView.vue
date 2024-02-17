<template>
  <Breadcrumb raid="riv" encounter="rww"/>

  <div class="parent">
    <PlateIcon v-for="value in wish.values" :symbol="value"/>
  </div>

  <div class="details">
    <p>Wish {{wish.id}}: <b>{{wish.name}}</b></p>
    <p>Effect: {{wish.effect}}</p>
  </div>

  <div class="keyboard">
    <div v-for="(_, index) in Array(14)" :class="selected(index+1)" @click="show(index+1)">{{index+1}}</div>
    <div @click="show(0)">B</div>
    <div @click="show(-1)">S</div>
  </div>
</template>

<script>
import Breadcrumb from "@/components/Breadcrumb";
import Storage from "@/data/storage";
import PlateIcon from "@/components/riv/PlateIcon.vue";
import {WishConfigurations} from "@/data/wishes";

const version = "v0.1"

export default {
  name: "WishingWallView",
  components: {PlateIcon, Breadcrumb},
  data: () => ({
    wish: WishConfigurations.BlankWall,
  }),
  created() {

  },
  computed: {

  },
  methods: {
    show(index) {
      if (index === -1) {
        this.wish = WishConfigurations.Symbols
        return
      }

      const wish = WishConfigurations.Walls[index]

      if (!wish) {
        this.wish = WishConfigurations.BlankWall
        console.error("unable to find wish at", index)
        return
      }

      this.wish = wish;
    },

    selected(index) {
      return index === this.wish.id ? 'selected' : ''
    }
  }
}
</script>

<style scoped>
.parent {
  margin: auto;
  max-width: 640px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: repeat(4, 100px);
  grid-column-gap: 10px;
  grid-row-gap: 0;
}

.details {
  margin: auto;
  max-width: 640px;
}

.details p {
  margin-bottom: .3em;
}

.keyboard {
  max-width: 640px;
  margin: 30px auto 0;
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 8px;
}

.keyboard div {
  width: 45px;
  height: 45px;
  border: 1px solid white;
  background: dimgrey;
  color: white;
  overflow: hidden;
  font-size: 1.3em;
  text-align: center;
  line-height: 45px;
  border-radius: 50%;
  user-select: none;
  cursor: pointer;
}

.keyboard div:hover {
  background: gray;
}

.keyboard div.selected {
  background: white;
  color: black;
}
</style>
