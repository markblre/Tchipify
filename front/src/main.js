import { createApp, markRaw } from 'vue'
import App from './App.vue'
import router from './router/index.js'

import "bootswatch/dist/vapor/bootstrap.css";
import "bootstrap/dist/js/bootstrap";
import '@fortawesome/fontawesome-free/css/all.min.css';
import 'vue-toast-notification/dist/theme-bootstrap.css'

import {createPinia} from "pinia"


/* Creating app */
const app = createApp(App)

/* Adding plugins to the app */
app.use(router)

/* Adding plugins to pinia*/
const pinia = createPinia()
pinia.use(() => ({
    $router: markRaw(router),
}))
/* Adding pinia to app */
app.use(pinia)

/* Mounting app */
app.mount('#app')