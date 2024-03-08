<script setup>
import { ref, computed, watch } from 'vue'
import {
  Iphone,
  // User,
  Lock,
  View,
  Hide,
  Message
} from '@element-plus/icons-vue'
import {
  pwdSendCodeService,
  pwdVerifyCodeService,
  pwdCodeService,
  pwdOldService
} from '@/api/pwdChange'
import navbarPC from '@/components/navbarPC.vue'
import { useRouter } from 'vue-router'
const router = useRouter()

// 用验证码修改密码时，区分是否已经验证了验证码，以区分展示的是关于验证码的页面还是输入要设置的新密码的页面
const isBefore = ref()
isBefore.value = false

// 监听isBefore的变化，如果变了，就让1.输入框清空，2.调用switchCard函数——为了让密码框初始时都不显示

watch(isBefore, () => {
  formModel.value = {
    id: '',
    code: '',
    oldPwd: '',
    newPwd: '',
    rePwd: ''
  }
  switchCard()
})

// 展示密码与否，图标显示
// true->展示->Hide->text
// false->不展示->View->password

// 控制图标
const showPwd = ref({
  pwd1: false,
  pwd2: false,
  pwd3: false
})

// 控制输入框
const showPwdType = ref({
  pwd1: computed(() => {
    return showPwd.value.pwd1 ? 'text' : 'password'
  }),
  pwd2: computed(() => {
    return showPwd.value.pwd2 ? 'text' : 'password'
  }),
  pwd3: computed(() => {
    return showPwd.value.pwd2 ? 'text' : 'password'
  })
})

// 切换展示情况的函数
const switchFun = {
  pwd1: () => {
    showPwd.value.pwd1 = !showPwd.value.pwd1
  },
  pwd2: () => {
    showPwd.value.pwd2 = !showPwd.value.pwd2
  },
  pwd3: () => {
    showPwd.value.pwd3 = !showPwd.value.pwd3
  }
}

// 切换卡片时要让密码展示情况都为false
const switchCard = () => {
  for (let i in showPwd.value) {
    showPwd.value[i] = false
  }
}
switchCard()

// 关于表单内容收集及表单验证
const formModel = ref({
  id: '',
  code: '',
  oldPwd: '',
  newPwd: '',
  rePwd: ''
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
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    {
      pattern: /^\d{11}$/,
      message: '手机号必须4位非空字符',
      trigger: 'blur'
    }
  ],
  oldPwd: [
    { required: true, message: '请输入旧密码', trigger: 'blur' },
    {
      pattern: /^\w{6,16}$/,
      message: '密码应为6至15位字母数字下划线',
      trigger: 'blur'
    }
  ],
  newPwd: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    {
      pattern: /^\w{6,16}$/,
      message: '密码应为6至15位字母数字下划线',
      trigger: 'blur'
    }
  ],
  rePwd: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== formModel.value.newPwd) {
          callback(new Error('两次输入密码不一致!'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

const form = ref(null)
const sendCode = async () => {
  console.log('hello')
  await form.value.validate()
  const res = await pwdSendCodeService({ id: formModel.value.id })
  if (res.status !== 200) {
    ElMessage.error(res.data.msg)
  } else {
    ElMessage.success('发送成功，请注意查收')
    time.value--
    timer.value = setInterval(() => {
      time.value--
      if (time.value === 0) {
        time.value = 60
        clearInterval(timer.value)
        timer.value = null
      }
    }, 1000)
  }
}

// 此处与发送验证码显示的文字有关
// 定时器、时间、显示的文字
const time = ref(60)
const timer = ref(null)
const sendCodeText = computed(function () {
  return time.value === 60 ? '发送验证码' : `${time.value}秒后重新发送`
})

// 此处是关于发送验证验证码请求的函数
const judgeCode = async () => {
  await form.value.validate()
  const res = await pwdVerifyCodeService({
    id: formModel.value.id,
    code: formModel.value.code
  })
  if (res.status !== 200) {
    ElMessage.error(res.data.msg)
  } else {
    ElMessage.error('身份验证成功')
    setTimeout(() => {
      isBefore.value = false
    }, 1000)
  }
}

// 此处是关于用验证码修改密码的请求
const changePwdCode = async () => {
  await form.value.validate()
  const res = await pwdCodeService({
    id: formModel.value.id,
    pwd: formModel.value.newPwd
  })
  if (res.status !== 200) {
    ElMessage.error(res.data.msg)
  } else {
    ElMessage.error('修改密码成功，即将跳转')
    setTimeout(() => {
      router.push('/login')
    }, 1000)
  }
}

// 此处是关于用旧密码修改密码的请求
const changePwdOld = async () => {
  await form.value.validate()
  const res = await pwdOldService({
    id: formModel.value.id,
    oldPwd: formModel.value.oldPwd,
    newPwd: formModel.value.newPwd
  })
  if (res.status !== 200) {
    ElMessage.error(res.data.msg)
  } else {
    ElMessage.error('修改密码成功，即将跳转')
    setTimeout(() => {
      router.push('/login')
    }, 1000)
  }
}

// 标签切换时的函数
const changeTab = () => {
  formModel.value = {
    id: '',
    code: '',
    oldPwd: '',
    newPwd: '',
    rePwd: ''
  }
  switchCard()
}
</script>
<template>
  <div class="pwdChange">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <div class="main">
        <el-tabs class="demo-tabs" stretch @tab-change="changeTab">
          <el-tab-pane>
            <template #label>
              <span class="tag"> 使用验证码修改 </span>
            </template>
            <div class="before" v-if="isBefore">
              <el-form
                class="elform"
                :model="formModel"
                :rules="rules"
                ref="form"
              >
                <el-form-item>
                  <h2>身份验证</h2>
                </el-form-item>
                <el-form-item prop="id" class="elinput">
                  <el-input
                    v-model="formModel.id"
                    placeholder="请输入学号"
                    :prefix-icon="Iphone"
                    ><template #append
                      ><el-button
                        style="padding: 7px"
                        @click="sendCode"
                        :disabled="time !== 60"
                        >{{ sendCodeText }}</el-button
                      ></template
                    ></el-input
                  >
                </el-form-item>
                <el-form-item prop="code" class="elinput">
                  <el-input
                    v-model="formModel.code"
                    placeholder="请输入短信验证码"
                    :prefix-icon="Message"
                  ></el-input>
                </el-form-item>
                <el-form-item>
                  <div class="btnGroup">
                    <el-button
                      type="primary"
                      size="large"
                      class="btn"
                      @click="judgeCode"
                      >确认</el-button
                    >
                    <el-button
                      type="primary"
                      size="large"
                      class="btn"
                      @click="router.back"
                      >取消</el-button
                    >
                  </div>
                </el-form-item>
              </el-form>
            </div>
            <div class="after" v-else>
              <el-form
                class="elform"
                :model="formModel"
                :rules="rules"
                ref="form"
              >
                <el-form-item>
                  <h2>修改密码</h2>
                </el-form-item>
                <el-form-item prop="newPwd" class="elinput">
                  <el-input
                    v-model="formModel.newPwd"
                    placeholder="请输入新密码"
                    :type="showPwdType.pwd1"
                    :prefix-icon="Lock"
                    ><template #append
                      ><el-icon v-if="!showPwd.pwd1" @click="switchFun.pwd1()">
                        <View /> </el-icon
                      ><el-icon v-else @click="switchFun.pwd1()">
                        <Hide /> </el-icon></template
                  ></el-input>
                </el-form-item>
                <el-form-item prop="rePwd" class="elinput">
                  <el-input
                    v-model="formModel.rePwd"
                    placeholder="请确认密码"
                    :type="showPwdType.pwd2"
                    :prefix-icon="Lock"
                    ><template #append
                      ><el-icon v-if="!showPwd.pwd2" @click="switchFun.pwd2()">
                        <View /> </el-icon
                      ><el-icon v-else @click="switchFun.pwd2()">
                        <Hide /> </el-icon></template
                  ></el-input>
                </el-form-item>
                <el-form-item>
                  <div class="btnGroup">
                    <el-button
                      type="primary"
                      size="large"
                      class="btn"
                      @click="changePwdCode"
                      >确认</el-button
                    >
                    <el-button
                      type="primary"
                      size="large"
                      class="btn"
                      @click="isBefore = true"
                      >取消</el-button
                    >
                  </div>
                </el-form-item>
              </el-form>
            </div>
          </el-tab-pane>
          <el-tab-pane>
            <template #label>
              <span class="tag"> 使用旧密码修改 </span>
            </template>
            <div class="before">
              <el-form
                class="elform"
                :model="formModel"
                :rules="rules"
                ref="form"
              >
                <el-form-item>
                  <h2>修改密码</h2>
                </el-form-item>
                <el-form-item prop="oldPwd" class="elinput">
                  <el-input
                    v-model="formModel.oldPwd"
                    placeholder="请输入旧密码"
                    :type="showPwdType.pwd1"
                    :prefix-icon="Lock"
                    ><template #append
                      ><el-icon v-if="!showPwd.pwd1" @click="switchFun.pwd1()">
                        <View /> </el-icon
                      ><el-icon v-else @click="switchFun.pwd1()">
                        <Hide /> </el-icon></template
                  ></el-input>
                </el-form-item>
                <el-form-item prop="newPwd" class="elinput">
                  <el-input
                    v-model="formModel.newPwd"
                    placeholder="请确认密码"
                    :type="showPwdType.pwd2"
                    :prefix-icon="Lock"
                    ><template #append
                      ><el-icon v-if="!showPwd.pwd2" @click="switchFun.pwd2()">
                        <View /> </el-icon
                      ><el-icon v-else @click="switchFun.pwd2()">
                        <Hide /> </el-icon></template
                  ></el-input>
                </el-form-item>
                <el-form-item prop="rePwd" class="elinput">
                  <el-input
                    v-model="formModel.rePwd"
                    placeholder="请确认密码"
                    :type="showPwdType.pwd3"
                    :prefix-icon="Lock"
                    ><template #append
                      ><el-icon v-if="!showPwd.pwd3" @click="switchFun.pwd3()">
                        <View /> </el-icon
                      ><el-icon v-else @click="switchFun.pwd3()">
                        <Hide /> </el-icon></template
                  ></el-input>
                </el-form-item>
                <el-form-item>
                  <div class="btnGroup">
                    <el-button
                      type="primary"
                      size="large"
                      class="btn"
                      @click="changePwdOld"
                      >确认</el-button
                    >
                    <el-button type="primary" size="large" class="btn"
                      >取消</el-button
                    >
                  </div>
                </el-form-item>
              </el-form>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>
<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}
.pwdChange {
  margin: -8px;
  // background: url('@/assets/bkg27.jpg') no-repeat;
  // background-size: 100% 100%;
  // background-attachment: fixed;
  .page {
    width: 80vw;
    margin: 0 auto;
    .main {
      padding-top: 20px; //这个属性是为了让标签栏高度变大，不然看着怪怪的
      border: 3px solid rgb(222, 154, 154);
      // background-color: rgba(255, 255, 255, 0.2);
      border-radius: 20px;
      .demo-tabs {
        .tag {
          color: rgb(55, 55, 55);
          margin-bottom: 15px;
        }
      }
      .elform {
        height: 50vh; // 这里是写死的，因为用%不起作用，尚不清楚原因
        margin: 0px auto;
        border: 1px solid rgb(255, 255, 255);
        border-radius: 20px;
        background-color: rgba(255, 255, 255, 0.3);
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        .elinput {
          margin: 0px auto;
        }
        h2 {
          color: rgb(0, 73, 182);
          font-size: 30px;
          letter-spacing: 20px;
          margin: 0px auto;
        }
        .btnGroup {
          width: 90%;
          margin: 0px auto;
          display: flex;
          justify-content: space-evenly;
          // margin-top: -30px;
          .btn {
            font-size: 2vw;
            border-radius: 4px;
            padding: 25px;
          }
        }
      }
    }
  }
}
@media (min-width: 600px) {
  .pwdChange {
    width: 100vw;
    height: 100vh;
    .main {
      width: 60vw;
      height: 74vh;
      min-height: 500px;
      min-width: 550px;
      margin: 0vh auto;
      .elform {
        width: 70%;
        .elinput {
          width: 70%;
        }
      }
      .demo-tabs {
        .tag {
          font-size: 20px;
          font-weight: 700;
        }
      }
    }
  }
}

@media (max-width: 600px) {
  .pwdChange {
    width: 100vw;
    height: 100vh;
    .main {
      width: 85vw;
      height: 74vh;
      min-width: 270px;
      margin: 0vh auto;
      border-color: black;
      .elform {
        width: 100%;
        .elinput {
          width: 90%;
        }
      }
      .demo-tabs {
        .tag {
          font-size: 16px;
          font-weight: 700;
        }
      }
    }
  }
}
</style>
