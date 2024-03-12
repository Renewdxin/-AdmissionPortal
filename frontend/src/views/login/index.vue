<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { Iphone, User, Lock, View, Hide } from '@element-plus/icons-vue'
import { userSignupService, userLoginService } from '@/api/user'
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()
const router = useRouter()
const isLogin = ref()
isLogin.value = false

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
</script>

<template>
  <div class="home">
    <div class="tl"></div>
    <div class="tr"></div>
    <div class="bl"></div>
    <div class="br"></div>
    <div class="main">
      <div class="left">
        <div class="img">
          <img src="@/assets/pic.png" alt="" />
        </div>
      </div>
      <div class="right">
        <div class="form">
          <div class="login" v-if="isLogin === true">
            <h1>登录</h1>
            <el-form
              class="elform"
              :model="formModel"
              :rules="rules"
              ref="form"
            >
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
            <el-text
              class="mx-1"
              type="primary"
              @click="router.push('/pwdChange')"
              >忘记密码</el-text
            >
          </div>
          <div class="register" v-if="isLogin === false">
            <h1>注册</h1>
            <el-form
              class="elform"
              :model="formModel"
              :rules="rules"
              ref="form"
            >
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
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0px;
  box-sizing: border-box;
}

.home {
  width: 92vw;
  height: 87vh;
  // border: 1px solid black;
  position: relative;
  padding-top: 1%;
  margin: 0 auto;
  margin-top: 50px;
  .tl,
  .tr,
  .bl,
  .br {
    position: absolute;
    width: 300px;
    height: 200px;
    background-color: rgb(221, 182, 247);
    z-index: -1;
  }
  .tl,
  .tr {
    top: 0px;
  }
  .bl,
  .br {
    bottom: 0px;
  }
  .tl,
  .bl {
    left: 0px;
  }
  .tr,
  .br {
    right: 0px;
  }
  .main {
    width: 98%;
    height: 98%;
    margin: 0px auto;
    position: relative;
    background-color: rgb(255, 255, 255);
    .left {
      width: 40vw;
      height: 40vw;
      .img {
        img {
          width: 100%;
        }
      }
    }

    .right {
      width: 55vw;
      height: 100%;
      position: absolute;
      right: 0px;
      top: 0px;

      .form {
        width: 55%;
        height: 70%;
        min-width: 300px;
        max-height: 530px;
        margin: 0% auto;
        margin-top: 9%;
        border-radius: 40px;
        background-color: rgba(242, 227, 249, 0.6);
        border: 5px solid rgb(255, 255, 255);
        box-shadow: rgba(226, 204, 247, 0.878) 0px 30px 30px -20px;

        .login,
        .register {
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
            .button1,
            .button2 {
              position: absolute;
              right: 0px;
              width: 100px;
              height: 50px;
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
        .register {
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
}

// @media (min-width: 640px) {
//   .home {
//     height: 100vh;
//   }
// }

// @media (max-width: 640px) {
//   .home {
//     width: 97vw;

//     .left {
//       width: 100%;
//       height: 97vw;
//       position: static;
//       margin-left: 0px;
//     }

//     .right {
//       position: static;
//       width: 100%;
//       height: 600px;

//       .form {
//         margin-top: 80px;
//       }
//     }
//   }
// }
</style>
