import request from '@/utils/request'

// 这里是关于注册、登录、报名
const userSignupService = (obj) => {
  return request.post('/auth/signup', obj)
}
const userLoginService = (obj) => {
  return request.post('/auth/login', obj)
}
const userApplyService = ({ obj }) => {
  return request.post(`/recruitment/job/${obj.id}/apply`, {
    name: obj.direction
  })
}

export { userSignupService, userLoginService, userApplyService }
