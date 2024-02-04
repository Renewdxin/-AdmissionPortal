<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import { userSignupService, userLoginService } from '@/api/user'
import { useTokenStore } from '@/stores/index'
const toeknStore = useTokenStore()
const route = useRoute()
const isLogin = ref()
isLogin.value = true
// const list = ref(['@/assets/3g.jpg', '@/assets/web.jpg', '@/assets/server.jpg', '@/assets/android.jpg', '@/assets/ios.jpg'])
// const imgList = computed(() => {
//     return list.value.map((item) => `@/assets/${item}.jpg`)
// })
// console.log(imgList.value);

// 关于表单内容收集及表单验证
const formModel = ref({
  tel: '',
  pwd: '',
  repwd: ''
})
const rules = {
  tel: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { min: 11, max: 11, message: '用户名必须是11位的字符', trigger: 'blur' }
  ],
  pwd: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码必须是6-15位的非空字符',
      trigger: 'blur'
    }
  ],
  repwd: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== formModel.value.pwd) {
          callback(new Error('两次输入密码不一致!'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}
const change = () => {
  isLogin.value = !isLogin.value
  formModel.value = {
    tel: '15592277810',
    pwd: '111111hh',
    repwd: '111111hh'
  }
}

const form = ref(null)
// 有对表单元素的内容的预检验，也就是validate函数，是异步操作
const signup = async () => {
  await form.value.validate()
  const res = await userSignupService({
    tel: formModel.value.tel,
    pwd: formModel.value.pwd
  })
  console.log(res)
  if (res.data.statue === 0) {
    // 注册成功
    ElMessage.success('注册成功，3秒后跳转')
    setTimeout(() => {
      isLogin.value = true
    }, 3000)
  } else {
    ElMessage.error(res.message)
  }
}
const login = async () => {
  await form.value.validate()
  const res = await userLoginService({
    tel: formModel.value.tel,
    pwd: formModel.value.pwd
  })
  console.log(res)
  if (res.data.statue === 0) {
    // 登录成功
    // 存token信息
    toeknStore.setToken(res.data.token)
    ElMessage.success('登录成功，3秒后跳转')
    setTimeout(() => {
      route.push('/user')
    }, 3000)
  } else {
    ElMessage.error(res.message)
  }
}
</script>

<template>
  <div class="home">
    <div class="left">
      <div class="circle">
        <div class="img">
          <div class="img1">
            <img src="@/assets/3g.jpg" alt="" />
          </div>
        </div>
        <div class="img">
          <div class="img1">
            <img src="@/assets/web.jpg" alt="" />
          </div>
        </div>
        <div class="img">
          <div class="img1">
            <img src="@/assets/server.jpg" alt="" />
          </div>
        </div>
        <div class="img">
          <div class="img1">
            <img src="@/assets/android.jpg" alt="" />
          </div>
        </div>
        <div class="img">
          <div class="img1">
            <img src="@/assets/ios.jpg" alt="" />
          </div>
        </div>
      </div>
    </div>
    <div class="right">
      <div class="form">
        <div class="login">
          <h1 v-if="isLogin === true">登录</h1>
          <h1 v-if="isLogin === false">注册</h1>

          <el-form class="elform" :model="formModel" :rules="rules" ref="form">
            <el-form-item prop="tel">
              <el-input
                class="elinput"
                dis
                v-model="formModel.tel"
                placeholder="请输入手机号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br /><br />
            <el-form-item prop="pwd">
              <el-input
                class="elinput"
                v-model="formModel.pwd"
                placeholder="请输入密码"
                type="password"
                :prefix-icon="Lock"
              ></el-input>
            </el-form-item>
            <br /><br />
            <el-form-item prop="repwd" v-if="isLogin === false">
              <el-input
                class="elinput"
                v-model="formModel.repwd"
                placeholder="请再次输入密码"
                type="password"
                size="large"
                :prefix-icon="Lock"
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="button" v-if="isLogin === true">
            <div class="button1" @click="login">登录</div>
            <div class="button2" @click="change">去注册</div>
          </div>
          <div class="button" v-if="isLogin === false">
            <div class="button1" @click="signup">注册</div>
            <div class="button2" @click="change">去登录</div>
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
  background: url('@/assets/bkg27.jpg') no-repeat;
  background-size: 100% 100%;
  //   background-color: antiquewhite;
  position: relative;

  .left {
    margin-left: 40px;
    width: 47%;
    // height: 100%;
    // border: 1px solid red;
    position: absolute;
    left: 0px;
    top: 0px;

    overflow: hidden;
    .circle {
      position: relative;
      width: 45vw;
      height: 45vw;
      border-radius: 50%;
      //   border: 1px solid red;
      margin-top: 50px;
      animation: rotation 20s linear infinite;
      // animation-play-state: paused; //控制动静

      &:hover {
        // transform: scale(0.5);不起作用
        animation-play-state: paused;
      }

      &:hover.circle img {
        animation-play-state: paused;
      }

      .img {
        position: absolute;
        width: 10vw;
        height: 10vw;
        top: 50%;
        left: 50%;
        margin-left: -5vw;
        margin-top: -5vw;
        border-radius: 50%;

        &:nth-child(1) {
          width: 14vw;
          height: 14vw;
          margin-left: -7vw;
          margin-top: -7vw;
        }

        .img1 {
          width: 100%;
          height: 100%;
          border-radius: 50%;
          //   background-color: pink;
        }

        @for $i from 2 through 5 {
          &:nth-child(#{$i}) {
            // background-color: aqua;
            transform-origin: center 21vw;
            transform: translate(0, -16vw) rotate(#{$i * 90-45}deg);

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
          // animation-play-state: paused; //控制动静
        }
      }
    }
  }

  .right {
    width: 52%;
    // height: 100%;
    // border: 1px solid black;
    position: absolute;
    right: 0px;
    top: 0px;

    .form {
      width: 65%;
      height: 70%;
      border: 1px solid red;
      margin: 0% auto;
      margin-top: 18%;
      border-radius: 40px;
      background-color: rgba(213, 226, 240, 0.6);
      border: 5px solid rgb(255, 255, 255);
      box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 30px 30px -20px;

      .login {
        .elform {
          width: 80%;
          margin: 0px auto;
          .elinput {
            border-color: transparent;
            border-bottom-color: rgb(75, 73, 73);
            border-bottom-width: 1px;
            outline: none;
            // background-color: rgb(255, 255, 255);
          }
        }
        h1 {
          letter-spacing: 30px;
          color: rgb(9, 110, 173);
          text-align: center;
        }
        .button {
          width: 220px;
          height: 50px;
          margin: 0px auto;
          position: relative;
          margin-bottom: 60px;

          // background-color: aquamarine;
          .button1,
          .button2 {
            position: absolute;
            right: 0px;
            width: 100px;
            height: 50px;
            // background-color: rgb(54, 54, 54);
            // border: 3px solid rgb(59, 59, 59);
            border-radius: 10px;
            color: azure;
            font-size: 18px;
            text-align: center;
            line-height: 40px;
            background: linear-gradient(
              45deg,
              rgb(16, 137, 211) 0%,
              rgb(18, 177, 209) 100%
            );
          }

          .button1 {
            left: 0px;
            // bottom:0px;
          }
        }
      }

      .login {
        h1 {
          margin-top: 60px;
          margin-bottom: 40px;
        }

        .button {
          margin-top: 1px;
        }
      }
    }
  }
}
@media (max-width: 640px) {
  .home {
    width: 97vw;
    .left {
      width: 100%;
      height: 97vw;
      position: static;
      margin-left: 0px;
      .circle {
        position: relative;
        width: 100%;
        height: 100%;
        margin-top: 10px;
        // border: 1px solid palegreen;
        .img {
          width: 24vw;
          height: 24vw;
          top: 50%;
          left: 50%;
          margin-left: -12vw;
          margin-top: -12vw;
          &:nth-child(1) {
            width: 32vw;
            height: 32vw;
            margin-left: -16vw;
            margin-top: -16vw;
          }
          @for $i from 2 through 5 {
            &:nth-child(#{$i}) {
              // background-color: aqua;
              transform-origin: center 47vw;
              transform: translate(0, -35vw) rotate(#{$i * 90-45}deg);

              .img1 {
                transform: rotate(#{45-$i * 90}deg);
              }
            }
          }
        }
      }
    }
    .right {
      position: static;
      width: 100%;
      height: 600px;
      .form {
        margin-top: 80px;
      }
    }
  }
}
</style>
