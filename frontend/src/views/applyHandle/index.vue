<script setup>
import navbarPC from '@/components/navbarPC.vue'
import router from '@/router'
// import { adminJobsService } from '@/api/admin'
import { ref } from 'vue'
const tableData = ref([])
const testRes = ref([])
// const adminJobs = async () => {
//   const res = await adminJobsService()
//   if (res.status === 200) {
//     tableData.value = res.data.info
//   } else {
//     ElMessage.error(res.data.msg)
//   }
// }
testRes.value = [
  {
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
  },
  {
    id: 2,
    name: '后台',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 2,
    name: '后台',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 2,
    name: '后台',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 2,
    name: '后台',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  },
  {
    id: 2,
    name: '后台',
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
      }
    ]
  },
  {
    id: 3,
    name: '安卓',
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
      }
    ]
  }
]
tableData.value = testRes.value
// testRes中apply_list是实际数组，tableData中的则是数组长度
tableData.value.forEach((item) => {
  item.apply_list = item.apply_list.length
})
console.log(tableData.value)
const applyDeatil = (index) => {
  console.log(index)
  router.push(`/applyList/${index}`)
}
</script>

<template>
  <div class="applyHandle">
    <div class="page">
      <navbarPC class="navbar"></navbarPC>
      <el-container class="content">
        <el-aside class="aside">
          <el-menu
            default-active="2"
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose"
            router
          >
            <el-sub-menu index="1">
              <template #title>
                <el-icon><location /></el-icon>
                <span>岗位分类</span>
              </template>
              <el-menu-item index="1-1">item one</el-menu-item>
              <el-menu-item index="1-2">item two</el-menu-item>
              <el-menu-item index="1-2">item two</el-menu-item>
              <el-menu-item index="1-2">item two</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>
        <el-main class="main">
          <el-table :data="tableData" stripe style="width: 100%">
            <el-table-column fixed prop="id" label="岗位编号" width="150" />
            <el-table-column fixed prop="name" label="岗位名称" width="150" />
            <el-table-column
              fixed
              prop="apply_list"
              label="申请人数"
              width="150"
            />
            <el-table-column fixed="right" label="操作" width="150">
              <template #default="scope">
                <el-button
                  link
                  type="primary"
                  @click="applyDeatil(scope.$index)"
                  >查看详情</el-button
                >
              </template>
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.applyHandle {
  margin: -8px;
  min-height: 100vh;
  background-color: rgb(242, 242, 242);
  .page {
    width: 80vw;
    margin: 0 auto;
    .content {
      min-height: 85vh;
      .aside {
        background-color: rgb(255, 255, 255);
        margin-right: 1vw;
      }
      .main {
        background-color: rgb(255, 255, 255);
      }
      .elheader {
        background-color: rgb(211, 211, 211);
        height: 65px;
        line-height: 65px;
        padding-left: 100px;
      }
    }
  }
}
</style>
