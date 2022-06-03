<template>
  <div class="keyboard-wrapper">
    <div class="keyboard">
      <div
          class="keyboard-key text-center"
          v-for="symbol in symbols"
          :key="symbol.id"
          @click="selectSymbol(symbol)"
      >
        <img class="symbol-display" :src="symbol.path" :class="highlight"/>
      </div>
    </div>
  </div>
</template>

<script>
import {symbols} from "@/data/symbols";

export default {
  name: "VodSymbolKeyboard",
  props: {
    only: []
  },
  data: () => ({
    selected: null
  }),
  computed: {
    symbols() {
      return symbols.filter(symbol => symbol.raid === 'vod')
    },
  },
  methods: {
    selectSymbol(symbol) {
      this.selected = symbol

      setTimeout(() => {
        this.selected = null
      }, 650)

      this.$emit('symbol:selected', symbol)
    },

    highlight(symbol) {
      if (!this.only.length) {
        return ''
      }

      return ''
    }
  }
}
</script>

<style scoped>
  .keyboard-wrapper {
    margin: 25px auto 0;
    max-width: 500px;
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