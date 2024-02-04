<script setup>
import { ref } from 'vue'
// import { infIsApplyService } from '@/api/inf'
import {
  User,
  Lock,
  Iphone,
  Location,
  OfficeBuilding
} from '@element-plus/icons-vue'
import { userApplyService } from '@/api/user'
import { infGetService, infUpdateService } from '@/api/inf'

// 是否编辑信息
const editInf = ref('')
editInf.value = 'false'
// 是否结束报名，如果结束，就展示面试信息，否则展示报名表单
const stopApply = ref('')
stopApply.value = 'false'

// 关于表单内容收集及表单验证
// 这个是修改信息表单中的内容
const formModel = ref({
  name: '',
  QQ: '',
  direction: '',
  class: '',
  gender: ''
})
const rules = {
  name: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { min: 11, max: 11, message: '用户名必须是11位的字符', trigger: 'blur' }
  ],
  QQ: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { min: 11, max: 11, message: '用户名必须是11位的字符', trigger: 'blur' }
  ],
  direction: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码必须是6-15位的非空字符',
      trigger: 'blur'
    }
  ],
  class: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
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
  ],
  gender: [{ required: true, message: '请再次输入密码', trigger: 'blur' }]
}
// 这个是报名表单中的内容
const formModel1 = ref({
  id: '',
  direction: ''
})
const rules1 = {
  id: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { min: 11, max: 11, message: '用户名必须是11位的字符', trigger: 'blur' }
  ]
}
// 这个是展示信息中的内容
const inf = ref({
  id: '04223147',
  name: '兰佳怡',
  tel: '15592277810',
  email: '15592277810@163.com',
  gender: '女'
})

const form = ref(null)
// 有对表单元素的内容的预检验，也就是validate函数，是异步操作
const apply = async () => {
  await form.value.validate()
  const res = await userApplyService(formModel.value)
  console.log(res)
  if (res.data.statue === 0) {
    // 报名成功
    ElMessage.success('报名成功')
  } else {
    ElMessage.error(res.message)
  }
}
const infUpdate = async () => {
  // 先修改后台信息
  const res = await infUpdateService(formModel.value)
  // 然后更新页面中的信息，重新渲染
  if (res.data.status === 0) {
    ElMessage.success('修改成功')
    const res1 = await infGetService(formModel.value.id)
    if (res1.data.status === 0) {
      inf.value = res1.data.inf
    } else {
      ElMessage.error(`${res1.data.message}`)
    }
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}
/* 功能：
1.查看信息：院系班级、QQ、姓名、性别、方向、（手机号、邮箱）、
2.查看面试状态
3.完善信息
4.报名

这个页面的访问要被限制，需要有身份证token


还需要的页面
1.修改密码&忘记密码（需要短信验证码）
2.管理员端
*/
</script>
<template>
  <div class="user">
    <div class="show">
      <div class="status">
        <!-- 报名部分只在未报名时显示 -->
        <div class="apply" v-if="stopApply === 'false'">
          <h1>报名</h1>
          <br /><br />
          <el-form
            class="elform"
            :model="formModel1"
            :rules="rules1"
            ref="form"
          >
            <el-form-item prop="id">
              <el-input
                class="elinput"
                v-model="formModel1.id"
                placeholder="请输入学号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br /><br />
            <el-form-item prop="direction" class="elselect">
              <el-select
                v-model="formModel1.direction"
                placeholder="请选择方向"
              >
                <el-option value="安卓" />
                <el-option value="iOS" />
                <el-option value="前端" />
                <el-option value="后台" />
                <el-option value="未定向" />
              </el-select>
            </el-form-item>
            <br />
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
            <el-form-item prop="id">
              <el-input
                class="elinput"
                v-model="formModel.id"
                placeholder="请输入学号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br />
            <el-form-item prop="name">
              <el-input
                class="elinput"
                v-model="formModel.name"
                placeholder="请输入姓名"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br />
            <el-form-item prop="tel">
              <el-input
                class="elinput"
                v-model="formModel.tel"
                placeholder="请输入手机号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br />
            <el-form-item prop="email">
              <el-input
                class="elinput"
                v-model="formModel.email"
                placeholder="请输入院系班级"
                :prefix-icon="Lock"
              ></el-input>
            </el-form-item>
            <br />
            <el-form-item prop="direction" class="elselect">
              <el-select v-model="formModel.direction" placeholder="请选择方向">
                <el-option value="安卓" />
                <el-option value="iOS" />
                <el-option value="前端" />
                <el-option value="后台" />
                <el-option value="未定向" />
              </el-select>
            </el-form-item>
            <br />
            <el-form-item prop="gender" class="elselect">
              <el-select v-model="formModel.gender" placeholder="请选择性别">
                <el-option value="男" />
                <el-option value="女" />
              </el-select>
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
                    <user />
                  </el-icon>
                  学号
                </div>
              </template>
              {{ inf.id }}
            </el-descriptions-item>
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
          </el-descriptions>
          <div class="buttonGroup">
            <div class="button" @click="editInf = 'true'">修改信息</div>
          </div>
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
  width: 100vw;
  .show {
    width: 75%;
    display: flex;
    justify-content: center;
    margin: 0% auto;
    // background-color: pink;
    .inf,
    .status {
      .edit,
      .noedit,
      .apply {
        width: 80%;
        margin: 0 auto;
        // background-color: rgb(179, 236, 236);
        .elform {
          width: 90%;
          margin: 0px auto;
          //   .elselect {
          //     width: 200px;
          //   }
        }
        .buttonGroup {
          display: flex;
          margin-top: 30px;
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
          margin-top: 10%;
          margin-bottom: 5%;
        }
      }
      .noedit {
        .cell-item {
          font-size: 2vh;
          min-width: 61px;
        }
      }
    }
    .inf,
    .status {
      width: 45%;
      height: 90%;
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
}
@media (min-width: 800px) {
  .user {
    height: 100vh;
    //   background: url('@/assets/bkg4.jpeg') no-repeat;
    background: url('@/assets/bkg27.jpg') no-repeat;
    background-size: 100% 100%;
    padding-top: 10vh;
    .show {
      height: 67vh;
      min-height: 450px;
    }
  }
}
@media (max-width: 800px) {
  .user {
    .show {
      width: 90vw;
      display: block;
      margin: 30px auto;
      //   background-color: pink;
      .inf,
      .status {
        width: 90%;
        height: 400px;
        margin: 0px auto;
        margin-bottom: 30px;
      }
    }
  }
}
</style>
