import request from '@/utils/request'

// 查看岗位总览
const recruitJobsService = () => {
  return request.get('/recruitment/jobs')
}
// 查看岗位详细信息
const recruitJobDetailService = (jobid) => {
  return request.get(`/recruitment/job/${jobid}`)
}
// 用户申请投递
const recruitApplyService = ({ jobid, userid, token }) => {
  return request.post(`/recruitment/job/${jobid}/apply/${userid}`, {
    token
  })
}

export { recruitJobsService, recruitJobDetailService, recruitApplyService }
