<template>
  <div class="keyboard-wrapper">
    <div class="keyboard">
      <div
          v-for="symbol in symbols"
          :key="symbol.id"
          @click="selectSymbol(symbol)"
          class="keyboard-key text-center"
      >
        <img :alt="symbol.id" class="symbol-display" :src="symbol.path"/>
      </div>
    </div>
  </div>
</template>

<script>
import {symbols} from "@/data/symbols";

export default {
  name: "KfSymbolKeyboard",
  props: {
    exclude: Array,
    only: Array
  },
  computed: {
    symbols() {
      const exclude = this.exclude || []
      const only = this.only || []

      const inRaid = symbols.filter(symbol => symbol.raid === 'kf')

      if (exclude.length) {
        return [...inRaid.filter(symbol => exclude.indexOf(symbol.id) === -1)]
      }

      if (only.length) {
        return [...inRaid.filter(symbol => only.indexOf(symbol.id) !== -1)]
      }

      return inRaid
    },
  },
  methods: {
    selectSymbol(symbol) {
      this.$emit('symbol:selected', symbol)
    }
  }
}
</script>

<style scoped>
  .keyboard-wrapper {
    margin: 25px auto 0;
    max-width: 500px;
    touch-action: manipulation;
  }

  .keyboard {
    display: grid;
    gap: 5px;
    grid-template-columns: repeat(6, auto);
  }

  .symbol-display {
    width: 50px;
  }
</style>