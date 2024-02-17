<script>
const data = [
  {id: 0, name: 'blank'},

  // Fish
  {id: 1, name: 'fish-left'},
  {id: 2, name: 'fish-right'},
  {id: 3, name: 'fish-circle'},
  {id: 4, name: 'fish-up'},

  // Snake
  {id: 5, name: 'snake-infinity'},
  {id: 6, name: 'snake-eight'},
  {id: 7, name: 'snake-two-head'},
  {id: 8, name: 'snake-cloud'},

  // Dragon
  {id: 9, name: 'dragon-left'},
  {id: 10, name: 'dragon-right'},
  {id: 11, name: 'dragon-up'},
  {id: 12, name: 'dragon-down'},

  // Bird
  {id: 13, name: 'bird-down'},
  {id: 14, name: 'bird-perched'},
  {id: 15, name: 'bird-left'},
  {id: 16, name: 'bird-behind'},
]

export default {
  name: "PlateIcon",
  props: {
    symbol: Number
  },
  computed: {
    styler() {
      return data[this.symbol].name
    }
  }
}
</script>

<template>
  <div class="wall-slot">
    <div class="mask">
      <img src="/assets/wishing-wall-symbols-small.png" :class="styler">
    </div>
    <div class="count">{{ symbol }}</div>
  </div>

</template>

<style scoped>
.wall-slot {
  display: inline-block;
}

.mask {
  position: relative;
  display: inline-block;
  width: 65px;
  height: 65px;
  overflow: hidden;
  border: 1px solid white;
  border-radius: 50%;
}

.count {
  width: 65px;
  height: 65px;
  text-align: center;
}

.mask img {
  position: absolute;
  width: 320px;
  height: auto;
}

.blank {
  top: 500px;
  left: 500px;
}

.fish-left {
  top: -87px;
  left: -12px;
}

.fish-right {
  top: -106px;
  left: -88px;
}

.fish-circle {
  top: -125px;
  left: -165px;
}

.fish-up {
  top: -145px;
  left: -240px;
}

.snake-infinity {
  top: -162px;
  left: -12px;
}

.snake-eight {
  top: -181px;
  left: -88px;
}

.snake-two-head {
  top: -200px;
  left: -165px;
}

.snake-cloud {
  top: -220px;
  left: -240px;
}

.dragon-left {
  top: -238px;
  left: -13px;
}

.dragon-right {
  top: -258px;
  left: -88px;
}

.dragon-up {
  top: -276px;
  left: -165px;
}

.dragon-down {
  top: -295px;
  left: -240px;
}

.bird-down {
  top: -13px;
  left: -12px;
}

.bird-perched {
  top: -31px;
  left: -88px;
}

.bird-left {
  top: -50px;
  left: -165px;
}

.bird-behind {
  top: -70px;
  left: -240px;
}
</style>
