<template>
  <div class="home">
    <div class="main">
      <div class="intro hover" @click="router.push('/intro')">
        <h1>实验室简介</h1>
        <p>know more about 3g</p>
      </div>
      <div class="login hover" @click="router.push('/login')">
        <h1>登录</h1>
        <p>LOGIN</p>
      </div>
      <div class="user hover" @click="router.push('/user')">
        <h1>个人信息</h1>
        <p>USER</p>
      </div>
      <div class="website hover" @click="goWebsite()">
        <h1>进入官网</h1>
        <p>WEBSITE</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import router from '@/router'
const goWebsite = () => {
  return (location.href = 'https://mobile.xupt.edu.cn/')
}
</script>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}

.home {
  margin: -8px;
  width: 100vw;
  height: 100vh;
  padding-top: 15vh;
  background: url('@/assets/bkg27.jpg') no-repeat;
  background-size: 100% 100%;
  background-attachment: fixed;

  .main {
    width: 67vw;
    height: 55vh;
    margin: 0px auto;
    display: flex;
    justify-content: space-around;

    .intro {
      background-color: rgba(244, 63, 94, 0.8);
    }

    .website {
      background-color: rgba(59, 130, 246, 0.8);
    }

    .login {
      background-color: rgba(34, 197, 94, 0.8);
    }

    .user {
      background-color: rgba(255, 223, 80, 0.8);
    }

    .intro,
    .website,
    .login,
    .user {
      width: 15vw;
      height: 50vh;
      margin-top: 2vh;
      border-radius: 3vh;
      transition: all 0.5s;

      p {
        font-size: 20px;
      }

      p,
      h1 {
        text-align: center;
        margin-top: 25%;
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

@media (max-width: 640px) {
  .home {
    .main {
      width: 65vw;
      height: 400px;
      flex-direction: column;
      justify-content: space-around;

      .intro,
      .website,
      .login,
      .user {
        margin: 0 auto;
        width: 88%;
        height: 23%;
        margin-top: 0vh;
        box-sizing: content-box;
        border-radius: 3vh;
        transition: all 0.5s;

        p {
          font-size: 18px;
        }

        p,
        h1 {
          text-align: center;
          margin-top: 4%;
          color: azure;
        }
      }
    }
  }
}
</style>
