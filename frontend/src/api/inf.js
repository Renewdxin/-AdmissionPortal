import request from '@/utils/request'

// 这里是修改后台数据
const infUpdateService = (id, { phoneNumber, email }, token) => {
  return request.put(`/profile/update/${id}`, { phoneNumber, email, token })
}
// 这里是更新拉取个人信息的数据
const infGetService = (id, token) => {
  return request.put(`/profile/Info/${id}`, { token })
}
// 删除用户
const infDeleteService = (id, token) => {
  return request.delete(`/profile/delete/${id}`, { token })
}
// 获取录取状态信息
const infStatusService = (id, token) => {
  return request.get(`/profile/status/${id}`, { token })
}

export { infUpdateService, infGetService, infDeleteService, infStatusService }
