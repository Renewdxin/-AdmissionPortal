import { createApp } from 'vue'
import pinia from './stores'

import 'sass-loader'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(pinia)
app.use(router)

app.mount('#app')
