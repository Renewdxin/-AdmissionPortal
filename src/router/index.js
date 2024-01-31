import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/home/index.vue'
import intro from '../views/intro/index.vue'
import login from '../views/login/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
      path: '/intro',
      name: 'intro',
      component: intro
    },
    {
      path: '/login',
      name: 'login',
      component: login
    }
  ]
})

export default router
