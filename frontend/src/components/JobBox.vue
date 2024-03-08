<script setup>
import router from '@/router'
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()
import { recruitApplyService } from '@/api/recruit'
const props = defineProps({
  name: String,
  id: Number
})
const apply = async () => {
  const res = await recruitApplyService(props.id, idStore.id, tokenStore.token)
  if (res.status === 200) {
    ElMessage.success('报名成功')
  } else {
    ElMessage.error(res.data.msg)
  }
}
const detail = () => {
  router.push(`/jobDetail?jobId=${props.id}&jobName=${props.name}`)
}
</script>

<template>
  <div class="jobBox">
    <div class="page">
      <span class="name">{{ name }}</span>
      <div class="btnGroup">
        <el-button class="btn" type="primary" plain @click="apply()"
          >报名</el-button
        >
        <el-button class="btn" type="primary" plain @click="detail()"
          >了解详情</el-button
        >
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.jobBox {
  height: 65px;
  margin-bottom: 10px;
  .page {
    width: 100%;
    height: 100%;
    background-color: rgb(255, 255, 255);
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: relative;
    .btnGroup {
      position: absolute;
      right: 5vw;
    }
    .name {
      position: absolute;
      left: 4vw;
    }
  }
}
</style>
