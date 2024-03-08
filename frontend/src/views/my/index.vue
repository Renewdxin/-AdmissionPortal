<script setup>
import { Location } from '@element-plus/icons-vue'
import navbarPC from '@/components/navbarPC.vue'
const handleOpen = (key, keyPath) => {
  console.log(key, keyPath)
}
const handleClose = (key, keyPath) => {
  console.log(key, keyPath)
}
</script>

<template>
  <div class="my">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <el-container class="content">
        <el-aside class="aside">
          <el-menu
            default-active="2"
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose"
            router
          >
            <el-sub-menu index="1">
              <template #title>
                <el-icon><location /></el-icon>
                <span>投递情况</span>
              </template>
              <el-menu-item-group title="Group One">
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
              </el-menu-item-group>
            </el-sub-menu>
            <el-sub-menu>
              <template #title>
                <el-icon><location /></el-icon>
                <span>个人信息</span>
              </template>
              <el-menu-item index="/user">查看资料</el-menu-item>
              <el-menu-item index="/pwdChange">修改密码</el-menu-item>
              <el-menu-item index="">退出登录</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>
        <el-main class="main">
          <router-view></router-view>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}
.my {
  margin: -8px;
  min-height: 100vh;
  background-color: rgb(242, 242, 242);
  .page {
    width: 80vw;
    margin: 0 auto;
    .content {
      min-height: 85vh;
      .aside {
        background-color: rgb(255, 255, 255);
        margin-right: 1vw;
      }
      .main {
        background-color: rgb(255, 255, 255);
      }
      .elheader {
        background-color: rgb(211, 211, 211);
        height: 65px;
        line-height: 65px;
        padding-left: 100px;
      }
    }
  }
}
</style>
