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
editInf.value = 'true'
// 是否结束报名，如果结束，就展示面试信息，否则展示报名表单
const stopApply = ref('')
stopApply.value = 'false'

// 关于表单内容收集及表单验证
// 这个是修改信息表单中的内容
const formModel = ref({
  id: '', //非必需
  name: '',
  // QQ: '',
  // direction: '',
  // class: '',
  gender: '',
  phoneNumber: '',
  email: '',
  birth: '',
  state: '' //非必需
})
const rules = {
  id: [
    // { required: true, message: '请输入学号', trigger: 'blur' },
    {
      pattern: /^\d{8}$/,
      message: '学号必须8位数字',
      trigger: 'blur'
    }
  ],
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    {
      pattern: /^[\u4e00-\u9fa5]{2,10}$/,
      message: '姓名必须是2到10位中文字符',
      trigger: 'blur'
    }
  ],
  // QQ: [
  //   { required: true, message: '请输入手机号', trigger: 'blur' },
  //   { min: 11, max: 11, message: '用户名必须是11位的字符', trigger: 'blur' }
  // ],
  // direction: [
  //   { required: true, message: '请输入密码', trigger: 'blur' },
  //   {
  //     pattern: /^\S{6,15}$/,
  //     message: '密码必须是6-15位的非空字符',
  //     trigger: 'blur'
  //   }
  // ],
  // class: [
  //   { required: true, message: '请再次输入密码', trigger: 'blur' },
  //   {
  //     validator: (rule, value, callback) => {
  //       if (value !== formModel.value.pwd) {
  //         callback(new Error('两次输入密码不一致!'))
  //       } else {
  //         callback()
  //       }
  //     },
  //     trigger: 'blur'
  //   }
  // ],
  gender: [{ required: true, message: '请选择性别', trigger: 'blur' }],
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
  ],
  birth: []
}

// 这个是报名表单中的内容
const formModel1 = ref({
  id: '',
  direction: ''
})
const rules1 = {
  id: [
    { required: true, message: '请输入学号', trigger: 'blur' },
    {
      pattern: /^\d{8}$/,
      message: '学号必须8位数字',
      trigger: 'blur'
    }
  ],
  direction: [{ required: true, message: '请选择方向', trigger: 'blur' }]
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
  const res = await userApplyService({ formModel1 })
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
// 所以钢刚进页面卜拉取数据，在提交一次信息之后才拉取
const infGet = async () => {
  const res = await infGetService(formModel.value.id)
  if (res.data.status === 0) {
    inf.value = res.data.inf
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}
const infUpdate = async () => {
  // 先修改后台信息
  const res = await infUpdateService(formModel.value)
  // 然后更新页面中的信息，重新渲染
  if (res.data.status === 0) {
    ElMessage.success('修改成功')
    infGet()
  } else {
    ElMessage.error(`${res.data.message}`)
  }
}

// const test = () => {
//   console.log(formModel.value)
//   console.log(formModel1.value)
// }
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
            <el-form-item prop="phoneNumber">
              <el-input
                class="elinput"
                v-model="formModel.phoneNumber"
                placeholder="请输入手机号"
                :prefix-icon="User"
              ></el-input>
            </el-form-item>
            <br />
            <el-form-item prop="email">
              <el-input
                class="elinput"
                v-model="formModel.email"
                placeholder="请输入邮箱"
                :prefix-icon="Lock"
              ></el-input>
            </el-form-item>
            <br />
            <!-- <el-form-item prop="direction" class="elselect">
              <el-select v-model="formModel.direction" placeholder="请选择方向">
                <el-option value="安卓" />
                <el-option value="iOS" />
                <el-option value="前端" />
                <el-option value="后台" />
                <el-option value="未定向" />
              </el-select>
            </el-form-item> -->
            <!-- <br /> -->
            <el-form-item prop="gender" class="elselect">
              <el-select v-model="formModel.gender" placeholder="请选择性别">
                <el-option value="男" />
                <el-option value="女" />
              </el-select>
            </el-form-item>
            <br />
            <el-form-item prop="birth" class="elselect">
              <el-date-picker
                v-model="formModel.birth"
                type="date"
                placeholder="请选择生日"
                clearable
              />
            </el-form-item>
          </el-form>
          <div class="buttonGroup">
            <div class="button" @click="infUpdate">确定</div>
            <div class="button" @click="editInf = 'false'">取消</div>
            <!-- <div class="button" @click="test">测试</div> -->
          </div>
        </div>
        <div class="noedit" v-else>
          <h1>个人信息</h1>
          <el-descriptions :column="1" border size="small">
            <!-- <el-descriptions-item class="cell-desc">
              <template #label>
                <div class="cell-item">
                  <el-icon :style="iconStyle">
                    <user />
                  </el-icon>
                  学号
                </div>
              </template>
              {{ inf.id }}
            </el-descriptions-item> -->
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
