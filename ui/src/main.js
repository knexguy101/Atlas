import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from "./router/index"
import store from './store/index'
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap.js"

window.store = store
createApp(App).use(router).use(store).mount('#app')