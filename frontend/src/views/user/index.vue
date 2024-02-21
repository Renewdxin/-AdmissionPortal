<script setup>
import { ref } from 'vue'
// import { infIsApplyService } from '@/api/inf'
import { User, Message } from '@element-plus/icons-vue'
import { userApplyService } from '@/api/user'
// import { ElMessage, ElMessageBox } from 'element-plus'
import { infGetService, infUpdateService, infDeleteService } from '@/api/inf'
import router from '@/router'
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()

// 是否编辑信息
const editInf = ref('')
editInf.value = 'true'
// 是否结束报名，如果结束，就展示面试信息，否则展示报名表单
const stopApply = ref('')
stopApply.value = 'false'

// 关于表单内容收集及表单验证
// 这个是修改信息表单中的内容
const formModel = ref({
  phoneNumber: '15592277810',
  email: '15592277810@163.com',
  name: '', //jobId
  direction: ''
})
const rules = {
  phoneNumber: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    {
      pattern: /^\d{11}$/,
      message: '手机号必须11位数字',
      trigger: 'blur'
    }
  ],
  email: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    {
      pattern: /^\d{11}$/,
      message: '手机号必须11位数字',
      trigger: 'blur'
    }
  ]
}
// 这个是展示信息中的内容
const inf = ref({
  // id: '04223147',
  name: '',
  tel: '',
  email: '',
  gender: '',
  birth: ''
})

const form = ref(null)
// 有对表单元素的内容的预检验，也就是validate函数，是异步操作
const apply = async () => {
  await form.value.validate()
  const res = await userApplyService(
    idStore.id,
    formModel.value.name,
    tokenStore.token
  )
  console.log(res)
  if (res.data.statue === 0) {
    // 报名成功
    ElMessage.success('报名成功')
  } else {
    ElMessage.error(res.message)
  }
}

// 拉取个人信息展示的数据，发请求
// 在刚进入页面时不知道学号id，因为在代码中，id是通过修改信息操作而得到
// 所以刚进页面卜拉取数据，在提交一次信息之后才拉取
const infGet = async () => {
  const res = await infGetService(idStore.id, tokenStore.token)
  if (res.data.status === 0) {
    inf.value = res.data.inf
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}
const infUpdate = async () => {
  // 先修改后台信息
  await form.value.validate()
  const res = await infUpdateService(idStore.id, formModel.value)
  // 然后更新页面中的信息，重新渲染
  if (res.data.status === 0) {
    ElMessage.success('修改成功')
    infGet()
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}
const deleteUser = async () => {
  ElMessageBox.confirm('确认删除用户？', '警告', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      const res = await infDeleteService(idStore.id)
      // 然后更新页面中的信息，重新渲染
      if (res.data.status === 0) {
        ElMessage.success('删除成功')
        router.push('/login')
      } else {
        ElMessage.error(`${res.data.message}`)
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消删除'
      })
    })
}
</script>
<template>
  <div class="user">
    <div class="status">
      <!-- 报名部分只在未报名时显示 -->
      <div class="apply" v-if="stopApply === 'false'">
        <h1>报名</h1>
        <el-form class="elform" :model="formModel" ref="form">
          <el-form-item prop="direction" class="elselect">
            <el-select v-model="formModel.direction" placeholder="请选择方向">
              <el-option value="安卓" />
              <el-option value="iOS" />
              <el-option value="前端" />
              <el-option value="后台" />
              <el-option value="未定向" />
            </el-select>
          </el-form-item>
        </el-form>
        <div class="buttonGroup">
          <div class="button" @click="apply">提交报名</div>
        </div>
      </div>
      <!-- 已报名情况下显示面试状态 -->
      <div class="process" v-else>已经通过一面</div>
    </div>
    <div class="inf">
      <!-- 编辑信息 -->
      <div class="edit" v-if="editInf === 'true'">
        <h1>修改信息</h1>
        <el-form class="elform" :model="formModel" :rules="rules" ref="form">
          <el-form-item prop="phoneNumber">
            <el-input
              class="elinput"
              v-model="formModel.phoneNumber"
              placeholder="请输入手机号"
              :prefix-icon="User"
            ></el-input>
          </el-form-item>
          <el-form-item prop="email">
            <el-input
              class="elinput"
              v-model="formModel.email"
              placeholder="请输入邮箱"
              :prefix-icon="Message"
            ></el-input>
          </el-form-item>
        </el-form>
        <div class="buttonGroup">
          <div class="button" @click="infUpdate">确定</div>
          <div class="button" @click="editInf = 'false'">取消</div>
        </div>
      </div>
      <div class="noedit" v-else>
        <h1>个人信息</h1>
        <el-descriptions :column="1" border size="small">
          <el-descriptions-item class="cell-desc">
            <template #label>
              <div class="cell-item">
                <el-icon :style="iconStyle">
                  <iphone />
                </el-icon>
                姓名
              </div>
            </template>
            {{ inf.name }}
          </el-descriptions-item>
          <el-descriptions-item class="cell-desc">
            <template #label>
              <div class="cell-item">
                <el-icon :style="iconStyle">
                  <location />
                </el-icon>
                手机号
              </div>
            </template>
            {{ inf.tel }}
          </el-descriptions-item>
          <el-descriptions-item class="cell-desc">
            <template #label>
              <div class="cell-item">
                <el-icon :style="iconStyle">
                  <office-building />
                </el-icon>
                邮箱
              </div>
            </template>
            {{ inf.email }}
          </el-descriptions-item>
          <el-descriptions-item class="cell-desc">
            <template #label>
              <div class="cell-item">
                <el-icon :style="iconStyle">
                  <office-building />
                </el-icon>
                性别
              </div>
            </template>
            {{ inf.gender }}
          </el-descriptions-item>
          <el-descriptions-item class="cell-desc">
            <template #label>
              <div class="cell-item">
                <el-icon :style="iconStyle">
                  <office-building />
                </el-icon>
                生日
              </div>
            </template>
            {{ inf.birth }}
          </el-descriptions-item>
        </el-descriptions>
        <div class="buttonGroup">
          <div class="button" @click="editInf = 'true'">修改信息</div>
          <div class="button" @click="deleteUser">删除用户</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
}

.user {
  margin: -8px;
  width: 100vw;
  background: url('@/assets/bkg27.jpg') no-repeat;
  background-size: 100% 100%;
  background-attachment: fixed;
  display: flex;
  .status,
  .inf {
    .edit,
    .noedit,
    .apply,
    .process {
      width: 80%;
      height: 90%;
      margin: 0 auto;

      // background-color: rgb(179, 236, 236);
      .elform {
        height: 72%;
        // .elform-item {
        //   height: 30%;
        // }
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        width: 90%;
        margin: 0px auto;
        // background-color: pink;
        //   .elselect {
        //     width: 200px;
        //   }
      }

      .buttonGroup {
        display: flex;
        margin-top: 25px;
        justify-content: space-evenly;

        //   background-color: pink;
        .button {
          width: 100px;
          height: 40px;
          margin: 0px auto;
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
      }

      h1 {
        font-size: 4vh;
        letter-spacing: 3vh;
        color: rgb(9, 110, 173);
        text-align: center;
        margin-top: 7%;
        margin-bottom: 4%;
      }
    }

    .noedit {
      .cell-item {
        font-size: 2vh;
        min-width: 61px;
      }
    }
  }

  .status,
  .inf {
    border: 5px solid rgb(255, 255, 255);
    border-radius: 40px;
    background-color: rgba(219, 230, 243, 0.8);
    box-shadow: rgba(133, 189, 215, 0.8784313725) 0px 30px 30px -20px;
    --color: #e1e1e1;
    background-image: linear-gradient(
        0deg,
        transparent 24%,
        var(--color) 25%,
        var(--color) 26%,
        transparent 27%,
        transparent 74%,
        var(--color) 75%,
        var(--color) 76%,
        transparent 77%,
        transparent
      ),
      linear-gradient(
        90deg,
        transparent 24%,
        var(--color) 25%,
        var(--color) 26%,
        transparent 27%,
        transparent 74%,
        var(--color) 75%,
        var(--color) 76%,
        transparent 77%,
        transparent
      );
    background-size: 55px 55px;
  }
}

@media (min-width: 600px) {
  .user {
    height: 100vh;
    // min-height: 600px;
    padding-top: 10vh;
    box-sizing: border-box;

    .inf,
    .status {
      width: 40%;
      height: 70%;
      margin: 0 auto;
      min-width: 400px;
    }
  }
}

@media (max-width: 600px) {
  .user {
    height: 100vh;
    padding-top: 12vh;
    box-sizing: border-box;

    .inf,
    .status {
      width: 85%;
      height: 62vh;
      margin: 0vh auto;
    }
  }
}
</style>
