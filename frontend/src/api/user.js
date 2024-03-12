import request from '@/utils/request'

// 这里是关于注册、登录
const userSignupService = ({ id, pwd: password, tel: phoneNumber }) => {
  return request.post('/auth/signup', { id, password, phoneNumber })
}
const userLoginService = ({ id, pwd: password }) => {
  return request.post('/auth/login', { id, password })
}

export { userSignupService, userLoginService }
