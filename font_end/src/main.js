import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import VueRouter from 'vue-router'
import router from './router'
import axios from 'axios'
import vueResource from 'vue-resource'


Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueRouter)
Vue.use(vueResource)
// axios.defaults.baseURL = 'http://localhost:8081/api';
Vue.prototype.$axios=axios

new Vue({
  el:'#app',
  render : h =>h(App),
  beforeCreate(){
      Vue.prototype.$bus=this
      // Vue.prototype.$axios=axios
  },
  router:router 
})