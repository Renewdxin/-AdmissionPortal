import request from '@/utils/request'

// 这里是是否报名判断
const infUpdateService = (obj) => {
  return request.put(`/profile/c/${obj.id}`, obj)
}
const infGetService = (id) => {
  return request.put(`/profile/Info/${id}`)
}

export { infUpdateService, infGetService }
