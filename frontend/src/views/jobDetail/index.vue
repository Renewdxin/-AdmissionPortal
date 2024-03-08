<script setup>
import { useRoute } from 'vue-router'
import { recruitJobDetailService, recruitApplyService } from '@/api/recruit'
import { ref } from 'vue'
import navbarPC from '@/components/navbarPC.vue'
const route = useRoute()
const jobId = route.query.jobId
const jobName = route.query.jobName
const detailInfo = ref('')
import { useTokenStore, useIdStore } from '@/stores/index'
const tokenStore = useTokenStore()
const idStore = useIdStore()
const props = defineProps({
  name: String,
  id: Number
})

const recruitJobDetail = async () => {
  const res = await recruitJobDetailService(jobId)
  if (res.status === 200) {
    detailInfo.value = res.data.info
  } else {
    ElMessage.error(res.data.msg)
  }
}
recruitJobDetail()
detailInfo.value = '这是岗位的详细信息'

const apply = async () => {
  const res = await recruitApplyService(props.id, idStore.id, tokenStore.token)
  if (res.status === 200) {
    ElMessage.success('报名成功')
  } else {
    ElMessage.error(res.data.msg)
  }
  ElMessage.success('报名成功')
}
</script>

<template>
  <div class="jobDetail">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <div class="card">
        <div class="header">
          <span class="icon">
            <svg
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                clip-rule="evenodd"
                d="M18 3a1 1 0 00-1.447-.894L8.763 6H5a3 3 0 000 6h.28l1.771 5.316A1 1 0 008 18h1a1 1 0 001-1v-4.382l6.553 3.276A1 1 0 0018 15V3z"
                fill-rule="evenodd"
              ></path>
            </svg>
          </span>
          <p class="alert">{{ jobName }}</p>
        </div>

        <p class="message">
          {{ detailInfo }}
        </p>

        <div class="actions">
          <span class="read" href="" @click="apply()"> 报名 </span>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.jobDetail {
  width: 90vw;
  margin: 0 auto;
  margin-top: -8px;
  box-sizing: border-box;
  .page {
    .card {
      width: 80vw;
      margin: 0 auto;
      border-width: 1px;
      border-color: rgba(219, 234, 254, 1);
      border-radius: 1rem;
      background-color: rgb(240, 240, 240);
      padding: 1rem;

      .header {
        display: flex;
        align-items: center;
        grid-gap: 1rem;
        gap: 1rem;
      }

      .icon {
        flex-shrink: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 9999px;
        background-color: rgb(95, 165, 252);
        padding: 0.5rem;
        color: rgba(255, 255, 255, 1);
      }

      .icon svg {
        height: 1rem;
        width: 1rem;
      }

      .alert {
        font-weight: 600;
        font-size: 20px;
        color: rgba(107, 114, 128, 1);
      }

      .message {
        font-size: 30px;
        margin-top: 1rem;
        color: rgba(107, 114, 128, 1);
      }

      .actions {
        margin-top: 1.5rem;
        display: flex;
        justify-content: center;
      }

      .read {
        display: inline-block;
        border-radius: 0.5rem;
        width: 70%;
        padding: 0.75rem 1.25rem;
        text-align: center;
        font-size: 0.875rem;
        line-height: 1.25rem;
        font-weight: 600;
      }

      .read {
        background-color: rgba(59, 130, 246, 1);
        color: rgba(255, 255, 255, 1);
      }
    }
  }
}
</style>
