<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()
import router from '@/router'
const name = route.name
console.log(name)
import { useTokenStore } from '@/stores/index'
const tokenStore = useTokenStore()
let navShow = ref('')
switch (name) {
  case 'jobDetail':
    navShow.value = '岗位详情'
    break
  case 'home':
    navShow.value = '主页'
    break
  case 'my':
    navShow.value = '个人中心'
    break
  case 'pwdChange':
    navShow.value = '修改密码'
    break
  case 'user':
    navShow.value = '用户信息'
    break
  case 'applyHandle':
    navShow.value = '职位列表'
    break
  case 'applyList':
    navShow.value = '申请列表'
    break

  default:
    break
}
</script>
<template>
  <div class="navbar">
    <div class="header">
      <span style="font-size: 30px; font-weight: 700">{{ navShow }}</span>
      <span style="position: absolute; right: 0px"
        ><span
          v-if="!tokenStore.token"
          style="font-size: 15px"
          @click="router.push('/login')"
          >登录</span
        >
        &nbsp;&nbsp;
        <span v-else style="font-size: 15px; font-weight: 700">
          你好，name </span
        >&nbsp;</span
      >
    </div>
  </div>
</template>

<style lang="scss" scoped>
.navbar {
  width: 100%;
  height: 7vh;
  margin-bottom: 4vh;
  background-color: rgb(255, 255, 255);
  .header {
    line-height: 7vh;
    border-bottom: 2px solid rgb(199, 199, 199);
    position: relative;
  }
}
</style>
