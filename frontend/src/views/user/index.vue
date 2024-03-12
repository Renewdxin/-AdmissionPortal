<script setup>
import { ref } from 'vue'
// import { infIsApplyService } from '@/api/inf'
// import { User, Message } from '@element-plus/icons-vue'
// import { ElMessage, ElMessageBox } from 'element-plus'
import { infGetService, infUpdateService, infDeleteService } from '@/api/inf'
import router from '@/router'
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()

// 是否编辑信息
const editInf = ref('')
editInf.value = 'true'
// 关于表单内容收集及表单验证
// 这个是修改信息表单中的内容
const ruleForm = ref({
  phoneNumber: '15592277810',
  email: '15592277810@163.com'
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
    { required: true, message: '请输入邮箱', trigger: 'blur' },
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

// 拉取个人信息展示的数据，发请求
const infGet = async () => {
  const res = await infGetService(idStore.id, tokenStore.token)
  if (res.status === 200) {
    inf.value = res.data.inf
    ruleForm.value.phoneNumber = inf.value.tel
    ruleForm.value.email = inf.value.email
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}
const infUpdate = async () => {
  // 先修改后台信息
  await form.value.validate()
  const res = await infUpdateService(idStore.id, ruleForm.value)
  // 然后更新页面中的信息，重新渲染
  if (res.status === 200) {
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
    <div class="inf">
      <!-- 编辑信息 -->
      <div class="edit" v-if="editInf === 'true'">
        <el-form
          ref="ruleFormRef"
          :model="ruleForm"
          :rules="rules"
          class="elform"
        >
          <el-form-item label="手机号" prop="phoneNumber">
            <el-input v-model="ruleForm.phoneNumber" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="ruleForm.email" />
          </el-form-item>
          <el-form-item class="btnGroup">
            <el-button type="primary" class="btn" @click="infUpdate"
              >确认</el-button
            >
            <el-button class="btn" @click="editInf = 'false'">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="noedit" v-else>
        <el-descriptions :column="1" border class="infShow">
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
        <el-form>
          <el-form-item class="btnGroup">
            <el-button type="primary" class="btn" @click="editInf = 'true'"
              >修改信息</el-button
            >
            <el-button class="btn" @click="deleteUser()">删除用户</el-button>
          </el-form-item>
        </el-form>
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
  .inf {
    width: 60%;
    min-width: 450px;
    .elform {
      height: 300px;
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
    }
    .infShow {
      margin-top: 30px;
      margin-bottom: 30px;
    }
    .btnGroup {
      // display: flex;
      // flex-direction: row;
      // justify-content: center;
      .btn {
        width: 80px;
        height: 40px;
      }
    }
  }
}
</style>
