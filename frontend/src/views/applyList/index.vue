<script setup>
import navbarPC from '@/components/navbarPC.vue'
import { adminHandleService } from '@/api/admin'
import { useRoute } from 'vue-router'
import { ref } from 'vue'
const jobName = ref('')
const route = useRoute()
const jobId = route.params.jobId
console.log(jobId)
const tableData = ref([])
const testRes = ref([])
testRes.value = {
  id: 1,
  name: '前端',
  apply_list: [
    {
      id: '',
      name: 'ljy',
      gender: '1',
      birth: '',
      email: '',
      phoneNumber: '',
      applyID: ''
    },
    {
      id: '',
      name: 'ljy',
      gender: '1',
      birth: '',
      email: '',
      phoneNumber: '',
      applyID: ''
    },
    {
      id: '',
      name: 'ljy',
      gender: '1',
      birth: '',
      email: '',
      phoneNumber: '',
      applyID: ''
    },
    {
      id: '',
      name: 'ljy',
      gender: '1',
      birth: '',
      email: '',
      phoneNumber: '',
      applyID: ''
    },
    {
      id: '',
      name: 'ljy',
      gender: '1',
      birth: '',
      email: '',
      phoneNumber: '',
      applyID: ''
    }
  ]
}
jobName.value = testRes.value.name
tableData.value = testRes.value.apply_list
const agree = async (index) => {
  const res = await adminHandleService(tableData.value[index].applyID)
  if (res.status === 200) {
    ElMessage.success('审批成功')
  } else {
    ElMessage.error(res.data.msg)
  }
  ElMessage.success('test:审批成功')
}
const refuse = async (index) => {
  const res = await adminHandleService(tableData.value[index].applyID)
  if (res.status === 200) {
    ElMessage.success('拒绝成功')
  } else {
    ElMessage.error(res.data.msg)
  }
  ElMessage.success('test:拒绝成功')
}
</script>
<template>
  <div class="applyList">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <div class="content">
        <el-table :data="tableData" stripe style="width: 100%">
          <el-table-column fixed prop="id" label="用户id" width="120" />
          <el-table-column fixed prop="name" label="姓名" width="120" />
          <el-table-column fixed prop="gender" label="性别" width="120" />
          <el-table-column fixed prop="birth" label="生日" width="120" />
          <el-table-column fixed prop="email" label="邮箱" width="120" />
          <el-table-column
            fixed
            prop="phoneNumber"
            label="手机号"
            width="120"
          />
          <el-table-column fixed prop="applyID" label="申请id" width="150" />
          <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
              <div class="btnGroup">
                <el-button link type="primary" @click="agree(scope.$index)"
                  >同意申请</el-button
                >
                <el-button link type="primary" @click="refuse(scope.$index)"
                  >拒绝申请</el-button
                >
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.applyList {
  margin: -8px;
  min-height: 100vh;
  background-color: rgb(242, 242, 242);
  .page {
    width: 80vw;
    margin: 0 auto;
    .content {
      min-height: 85vh;
      .btnGroup {
        display: flex;
        white-space: nowrap;
      }
    }
  }
}
</style>
