import request from '@/utils/request'

// 管理员仪表板或主页
const adminDashboardService = () => {
  return request.get('/admin/dashboard')
}
// 查看所有职位发布
const adminJobsService = () => {
  return request.get('/admin/jobs')
}
// 查看职位详情
const adminJobDetailService = (jobid) => {
  return request.get(`/admin/job/${jobid}`)
}
// 查看职位申请
const adminApplicationsService = (jobid) => {
  return request.get(`/admin/applications/${jobid}`)
}
// 审批或拒绝职位申请 + 自动发送录取短信
const adminHandleService = (appid) => {
  return request.get(`/admin/application/${appid}`)
}

export {
  adminDashboardService,
  adminJobsService,
  adminJobDetailService,
  adminApplicationsService,
  adminHandleService
}
