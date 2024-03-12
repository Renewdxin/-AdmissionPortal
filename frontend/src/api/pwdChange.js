import request from '@/utils/request'

// 这里是关于发送验证码、验证验证码、修改密码（两种）
const pwdSendCodeService = ({ id }) => {
  return request.post('/auth/code/send', { id })
}
const pwdVerifyCodeService = ({ id, code }) => {
  return request.post('/auth/code/verify', { id, code })
}
const pwdCodeService = ({ id, pwd: password }) => {
  return request.post('/auth/password/change/code', { id, password })
}
const pwdOldService = ({ id, oldPwd: oldPassword, newPwd: newPassword }) => {
  return request.post('/auth/password/change/pwd', {
    id,
    oldPassword,
    newPassword
  })
}

export {
  pwdSendCodeService,
  pwdVerifyCodeService,
  pwdCodeService,
  pwdOldService
}
