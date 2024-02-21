import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/home/index.vue'
import intro from '../views/intro/index.vue'
import login from '../views/login/index.vue'
import user from '../views/user/index.vue'
import pwdChange from '../views/pwdChange/index.vue'
import test from '../views/test/index.vue'
import notFound from '../views/notFound/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: 'home'
    },
    {
      path: '/home',
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
    },
    {
      path: '/user',
      name: 'user',
      component: user
    },
    {
      path: '/pwdChange',
      name: 'pwdChange',
      component: pwdChange
    },
    {
      path: '/test',
      name: 'test',
      component: test
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'notFound',
      component: notFound
    }
  ]
})

export default router
