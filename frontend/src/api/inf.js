import request from '@/utils/request'

// 这里是修改后台数据
const infUpdateService = (obj) => {
  return request.put(`/profile/c/${obj.id}`, obj)
}
// 这里是更新拉取个人信息的数据
const infGetService = (id) => {
  return request.put(`/profile/Info/${id}`)
}

export { infUpdateService, infGetService }
