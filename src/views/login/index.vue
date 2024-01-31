<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()
const inf = route.query.isLogin
const isLogin = ref()
isLogin.value = inf
const list = ['3g', 'web', 'android', 'server', 'ios']
</script>

<template>
  <div class="home">
    <div class="left">
      <div class="circle">
        <div class="img" v-for="item in list" :key="item">
          <div class="img1">
            <img src="`@/assets/web.jpg`" alt="" />
          </div>
        </div>
      </div>
    </div>
    <div class="right">
      <div class="form">
        <div class="signup" v-if="false">
          <h1>报名</h1>
          <input type="text" placeholder="姓名" />
          <input type="text" placeholder="电话" />
          <input type="text" placeholder="QQ" />
          <input type="text" placeholder="学院" />
          <input type="text" placeholder="专业" />
          <select name="" id="">
            <option value="">前端</option>
            <option value="">后台</option>
            <option value="">安卓</option>
            <option value="">iOS</option>
            <option value="">未定向</option>
          </select>
          <div class="button">
            <div class="button1">报名</div>
            <div class="button2">去登录</div>
          </div>
        </div>
        <div class="login" v-if="true">
          <h1>登录</h1>
          <input type="text" placeholder="姓名" />
          <input type="text" placeholder="电话" />
          <div class="button">
            <div class="button1">登录</div>
            <div class="button2">去注册</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0px;
  box-sizing: border-box;
}
@keyframes rotation {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@keyframes border {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
.home {
  margin: -8px;
  height: 100vh;
  width: 100vw;
  background-color: antiquewhite;
  position: relative;
  .left {
    width: 60%;
    height: 100%;
    border: 1px solid red;
    position: absolute;
    left: 0px;
    top: 0px;
    // overflow: hidden;
    .circle {
      position: relative;
      width: 40vw;
      height: 40vw;
      border-radius: 50%;
      border: 1px solid red;
      margin-top: 100px;
      animation: rotation 20s linear infinite;
      animation-play-state: paused; //控制动静
      &:hover {
        // transform: scale(0.5);不起作用
        animation-play-state: paused;
      }
      &:hover.circle img {
        animation-play-state: paused;
      }
      .img {
        position: absolute;
        width: 8vw;
        height: 8vw;
        top: 50%;
        left: 50%;
        margin-left: -4vw;
        margin-top: -4vw;
        border-radius: 50%;
        &:nth-child(1) {
          width: 12vw;
          height: 12vw;
          margin-left: -6vw;
          margin-top: -6vw;
        }
        .img1 {
          width: 100%;
          height: 100%;
          border-radius: 50%;
          // background-color: pink;
        }
        @for $i from 2 through 5 {
          &:nth-child(#{$i}) {
            // background-color: aqua;
            transform-origin: center 18vw;
            transform: translate(0, -14vw) rotate(#{$i * 90-45}deg);
            .img1 {
              transform: rotate(#{45-$i * 90}deg);
            }
          }
        }
        img {
          width: 100%;
          height: 100%;
          border-radius: 50%;
          animation: rotation 20s linear infinite reverse;
          animation-play-state: paused; //控制动静
        }
      }
    }
  }
  .right {
    width: 40%;
    height: 100%;
    border: 1px solid black;
    position: absolute;
    right: 0px;
    top: 0px;
    .form {
      width: 70%;
      height: 70%;
      border: 1px solid red;
      margin-top: 100px;
      margin-left: 50px;
      border-radius: 40px;
      .signup,
      .login {
        h1 {
          letter-spacing: 30px;
          color: rgb(84, 84, 84);
          text-align: center;
          margin-top: 30px;
        }
        input,
        select {
          display: block;
          margin: 20px auto;
          width: 250px;
          height: 30px;
          outline: none;
        }
        .button {
          width: 250px;
          height: 50px;
          margin: 0px auto;
          position: relative;
          // background-color: aquamarine;
          .button1,
          .button2 {
            position: absolute;
            right: 0px;
            width: 100px;
            height: 50px;
            background-color: rgb(54, 54, 54);
            border-radius: 10px;
            border: 3px solid rgb(59, 59, 59);
            color: azure;
            font-size: 18px;
            text-align: center;
            line-height: 40px;
          }
          .button1 {
            left: 0px;
            // bottom:0px;
          }
        }
      }
      .login {
        h1 {
          margin-top: 70px;
          margin-bottom: 70px;
        }
        .button {
          margin-top: 40px;
        }
      }
    }
  }
}
</style>
