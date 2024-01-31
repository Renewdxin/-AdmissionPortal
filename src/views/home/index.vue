<template>
  <div class="home">
    <div class="main">
      <div class="intro hover" @click="router.push('/intro')">
        <h1>实验室简介</h1>
        <p>know more about 3g</p>
      </div>
      <div class="register hover" @click="router.push('/login?isLogin=false')">
        <h1>报名</h1>
        <p>REGISTER</p>
      </div>
      <div class="login hover" @click="router.push('/login?isLogin=true')">
        <h1>登录</h1>
        <p>LOGIN</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import router from '@/router'
</script>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}
.home {
  width: 100vw;
  height: 100vh;
  padding-top: 15vh;
  background: url('@/assets/bkg1.webp') no-repeat;
  background-size: 100% 100%;
  .main {
    width: 67vw;
    height: 55vh;
    margin: 0px auto;
    background-color: transparent;
    display: flex;
    justify-content: space-around;
    .intro {
      background-color: rgb(244, 63, 94);
    }
    .register {
      background-color: rgb(59, 130, 246);
    }
    .login {
      background-color: rgb(34, 197, 94);
    }
    .intro,
    .register,
    .login {
      width: 20vw;
      height: 50vh;
      //   border: 1px solid black;
      margin-top: 2vh;
      border-radius: 3vh;
      transition: all 0.5s;
      p {
        font-size: 20px;
      }
      p,
      h1 {
        text-align: center;
        margin-top: 20%;
        color: azure;
      }
    }
  }
  .main:hover.main .hover:not(:hover) {
    filter: blur(10px);
  }
  .hover:hover {
    transform: scale(1.1, 1.1);
  }
}
</style>
