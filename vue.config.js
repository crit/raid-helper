const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    allowedHosts: "all",
    client: {
      webSocketURL: 'auto://0.0.0.0:0/ws'
    }
  },
  transpileDependencies: true,
  publicPath: ''
})
