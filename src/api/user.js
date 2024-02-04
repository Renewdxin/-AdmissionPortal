import request from '@/utils/request'

// 这里是关于注册、登录、报名
const userSignupService = ({ tel, pwd }) => {
  return request.post('/auth/signup', { tel, pwd })
}
const userLoginService = ({ tel, pwd }) => {
  return request.post('/auth/login', { tel, pwd })
}
const userApplyService = ({ tel, pwd }) => {
  return request.post('/auth/login', { tel, pwd })
}

export { userSignupService, userLoginService, userApplyService }
