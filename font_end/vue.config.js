// const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// })

module.exports = {
  pages: {
    index: {
      // page 的入口
      entry: 'src/main.js',
    },
  },
  // lintOnSave:false,
  //开启代理服务器 方法一
  // devServer: {
  //   proxy: 'http://localhost:5000'
  // }
  //开启代理服务器 方法二
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8081',
        pathRewrite:{'^/api':''}, 
        ws: false,     //用于支持websocket 
        changeOrigin: false  //用于控制请求头中的host值
      }
      // '/demo': {
      //   target: 'http://localhost:5001',
      //   pathRewrite:{'^/demp':''}, 
      //   ws: true,     //用于支持websocket 
      //   changeOrigin: false  //用于控制请求头中的host值
      // },
    }
  }
}