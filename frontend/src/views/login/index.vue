<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  Iphone,
  User,
  Lock,
  View,
  Hide,
  Message
} from '@element-plus/icons-vue'
import { userSignupService, userLoginService } from '@/api/user'
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()
const router = useRouter()
const isLogin = ref()
isLogin.value = false

// 忘记密码操作
const isForgetPwd = ref()
isForgetPwd.value = true

// 修改密码
const changePwd = ref()
changePwd.value = false

// 展示密码与否，图标显示
// true->展示->Hide->text
// false->不展示->View->password
const showPwd = ref()
showPwd.value = true
// 展示密码与否，输入框类型
const showPwdType = computed(() => {
  return showPwd.value ? 'password' : 'text'
})
// 展示密码与否，函数
const changePwdShow = () => {
  showPwd.value = !showPwd.value
}

// 确认密码相对应的
const showRepwd = ref()
showRepwd.value = true
// 展示密码与否，输入框类型
const showRepwdType = computed(() => {
  return showRepwd.value ? 'password' : 'text'
})
// 展示密码与否，函数
const changeRepwdShow = () => {
  showRepwd.value = !showRepwd.value
}

// 关于表单内容收集及表单验证
const formModel = ref({
  id: '04223147',
  tel: '15592277810',
  pwd: '111111',
  repwd: '111111'
})
const rules = {
  id: [
    { required: true, message: '请输入学号', trigger: 'blur' },
    {
      pattern: /^\d{8}$/,
      message: '手机号必须8位数字',
      trigger: 'blur'
    }
  ],
  tel: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    {
      pattern: /^\d{11}$/,
      message: '手机号必须11位数字',
      trigger: 'blur'
    }
  ],
  pwd: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\w{6,16}$/,
      message: '密码应为6至15位字母数字下划线',
      trigger: 'blur'
    }
  ],
  repwd: [
    { required: true, message: '请确认密码', trigger: 'blur' },
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
  showPwd.value = true
  showRepwd.value = true
  formModel.value = {
    id: '04223147',
    tel: '15592277810',
    pwd: '111111',
    repwd: '111111'
  }
}

const form = ref(null)
// 有对表单元素的内容的预检验，也就是validate函数，是异步操作
const signup = async () => {
  await form.value.validate()
  const res = await userSignupService(formModel.value)
  console.log(res)
  if (res.data.status === 0) {
    // 注册成功
    ElMessage.success('注册成功，1秒后跳转')
    setTimeout(() => {
      change()
    }, 1000)
  } else {
    ElMessage.error(res.data.message)
  }
}
const login = async () => {
  await form.value.validate()
  const res = await userLoginService(formModel.value)
  console.log(res)
  if (res.data.status === 0) {
    // 登录成功
    // 存token信息
    tokenStore.setToken(res.data.token)
    idStore.setId(formModel.value.id)
    // subTime()
    ElMessage.success('登录成功，3秒后跳转')
    setTimeout(() => {
      router.push('/user')
    }, 3000)
  } else {
    ElMessage.error(res.data.message)
  }
}
const forgetPwd = () => {}
// 倒计时时间
// let time = ref(3)
// let timer = null
// const subTime = () => {
//   timer = setInterval(() => {
//     time.value--
//     console.log(time.value)
//     if (time.value === 0) {
//       time.value = 3
//       clearInterval(timer)
//       timer = null
//     }
//   }, 1000)
// }
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
      <div class="forgetPwd" v-if="isForgetPwd === true">
        <div class="forget">
          <h1>忘记密码</h1>
          <el-form class="elform" :model="formModel" :rules="rules" ref="form">
            <el-form-item prop="tel">
              <el-input
                class="elinput"
                v-model="formModel.tel"
                placeholder="请输入手机号"
                :type="showPwdType"
                :prefix-icon="Iphone"
                ><template #append>发送验证码</template></el-input
              >
            </el-form-item>
            <el-form-item prop="code">
              <el-input
                class="elinput"
                v-model="formModel.tel"
                placeholder="请输入短信验证码"
                :type="showPwdType"
                :prefix-icon="Message"
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="button">
            <div class="button1" @click="forgetPwd">确认</div>
            <div class="button2" @click="isForgetPwd = false">取消</div>
          </div>
        </div>
      </div>
      <div class="form" v-if="!isForgetPwd">
        <div class="login" v-if="isLogin === true">
          <h1>登录</h1>
          <el-form class="elform" :model="formModel" :rules="rules" ref="form">
            <el-form-item prop="id">
              <el-input
                class="elinput"
                dis
                v-model="formModel.id"
                placeholder="请输入学号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <el-form-item prop="pwd">
              <el-input
                class="elinput"
                v-model="formModel.pwd"
                placeholder="请输入密码"
                :type="showPwdType"
                :prefix-icon="Lock"
                ><template #append
                  ><el-icon v-if="showPwd" @click="changePwdShow">
                    <View /> </el-icon
                  ><el-icon v-else @click="changePwdShow">
                    <Hide /> </el-icon></template
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="button">
            <div class="button1" @click="login">登录</div>
            <div class="button2" @click="change">去注册</div>
          </div>
          <el-text class="mx-1" type="primary" @click="isForgetPwd = true"
            >忘记密码</el-text
          >
        </div>
        <div class="register" v-if="isLogin === false">
          <h1>注册</h1>
          <el-form class="elform" :model="formModel" :rules="rules" ref="form">
            <el-form-item prop="tel">
              <el-input
                class="elinput"
                dis
                v-model="formModel.id"
                placeholder="请输入学号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <el-form-item prop="tel">
              <el-input
                class="elinput"
                dis
                v-model="formModel.tel"
                placeholder="请输入手机号"
                :prefix-icon="Iphone"
              ></el-input>
            </el-form-item>
            <el-form-item prop="pwd">
              <el-input
                class="elinput"
                v-model="formModel.pwd"
                placeholder="请输入密码"
                :type="showPwdType"
                :prefix-icon="Lock"
                ><template #append
                  ><el-icon v-if="showPwd" @click="changePwdShow">
                    <View /> </el-icon
                  ><el-icon v-else @click="changePwdShow">
                    <Hide /> </el-icon></template
              ></el-input>
            </el-form-item>
            <el-form-item prop="repwd">
              <el-input
                class="elinput"
                v-model="formModel.repwd"
                placeholder="请确认密码"
                :type="showRepwdType"
                :prefix-icon="Lock"
                ><template #append
                  ><el-icon v-if="showRepwd" @click="changeRepwdShow">
                    <View /> </el-icon
                  ><el-icon v-else @click="changeRepwdShow">
                    <Hide /> </el-icon></template
              ></el-input>
            </el-form-item>
          </el-form>
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
  width: 100vw;
  background: url('@/assets/bkg27.jpg') no-repeat;
  background-size: 100% 100%;
  background-attachment: fixed;
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
    height: 100%;
    position: absolute;
    right: 0px;
    top: 0px;

    .form,
    .forgetPwd {
      width: 65%;
      height: 66%;
      min-width: 300px;
      max-height: 530px;
      margin: 0% auto;
      margin-top: 18%;
      border-radius: 40px;
      background-color: rgba(213, 226, 240, 0.6);
      // background-color: rgb(215, 176, 183);
      border: 5px solid rgb(255, 255, 255);
      box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 30px 30px -20px;

      .login,
      .register,
      .forget {
        // background-color: bisque;
        height: 70%;
        position: relative;

        .elform {
          width: 80%;
          height: 70%;
          margin: 0px auto;
          // background-color: rgb(210, 182, 236);
          display: flex;
          flex-direction: column;
          justify-content: space-evenly;
        }

        h1 {
          letter-spacing: 30px;
          color: rgb(9, 110, 173);
          text-align: center;
        }
        .mx-1 {
          position: absolute;
          right: 25px;
          margin-top: 15px;
        }

        .button {
          width: 220px;
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

      .login,
      .register,
      .forget {
        h1 {
          margin-top: 60px;
          // margin-bottom: 10px;
        }

        .button {
          margin-top: 1px;
        }
      }

      .register {
        height: 100%;

        h1 {
          margin-top: 16px;
          margin-bottom: 5px;
        }
      }
    }
  }
}

@media (min-width: 640px) {
  .home {
    height: 100vh;
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

      .form,
      .forgetPwd {
        margin-top: 80px;
      }
    }
  }
}
</style>
