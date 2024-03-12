import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/home/index.vue'
import login from '../views/login/index.vue'
import user from '../views/user/index.vue'
import my from '../views/my/index.vue'
import jobDetail from '../views/jobDetail/index.vue'
import pwdChange from '../views/pwdChange/index.vue'
import applyHandle from '../views/applyHandle/index.vue'
import applyList from '../views/applyList/index.vue'
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
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/my',
      name: 'my',
      component: my,
      children: [
        {
          path: '/user',
          name: 'user',
          component: user
        }
      ]
    },
    {
      path: '/pwdChange',
      name: 'pwdChange',
      component: pwdChange
    },
    {
      path: '/jobDetail',
      name: 'jobDetail',
      component: jobDetail
    },
    {
      path: '/applyHandle',
      name: 'applyHandle',
      component: applyHandle
    },
    {
      path: '/applyList/:jobId',
      name: 'applyList',
      component: applyList
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
