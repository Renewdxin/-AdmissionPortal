<script setup>
import JobBox from '@/components/JobBox.vue'
import navbarPC from '@/components/navbarPC.vue'
import { ref } from 'vue'
// import router from '@/router'
import { recruitJobsService } from '@/api/recruit'

const jobTitleList = ref([])
jobTitleList.value = [
  {
    name: '前端',
    id: 1
  },
  {
    name: '后台',
    id: 2
  },
  {
    name: '安卓',
    id: 3
  },
  {
    name: 'iOS',
    id: 4
  }
]
const recruitJobs = async () => {
  const res = await recruitJobsService()
  if (res.status === 200) {
    jobTitleList.value = res.data.title
  } else {
    ElMessage.error(res.data.msg)
  }
}
recruitJobs()
</script>

<template>
  <div class="home">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <el-container class="main">
        <el-aside class="aside">
          <el-menu
            default-active="2"
            @open="handleOpen"
            @close="handleClose"
            router
          >
            <el-sub-menu>
              <template #title>
                <el-icon><location /></el-icon>
                <span>岗位总览</span>
              </template>
              <el-menu-item index="1">item one</el-menu-item>
              <el-menu-item index="2">item two</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>
        <el-main class="content">
          <div class="jobBox">
            <JobBox
              v-for="item in jobTitleList"
              :key="item.id"
              :name="item.name"
              :id="item.id"
            >
            </JobBox>
          </div>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
}
.home {
  margin: -8px;
  min-height: 100vh;
  background-color: rgb(242, 242, 242);
  padding-bottom: 2vh;
  .page {
    width: 90vw;
    margin: 0 auto;
    .main {
      margin-top: 4vh;
      min-height: 85vh;
      .aside {
        background-color: rgb(255, 255, 255);
        margin-right: 1vw;
      }
    }
  }
}
</style>
