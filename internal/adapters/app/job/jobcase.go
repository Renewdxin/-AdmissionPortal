package job

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type JobcaseAdapter struct {
	core job.JobsCorePorts
	dao  database.JobDaoPorts
}

func NewAdapter(core job.JobsCorePorts, dao database.JobDaoPorts) JobcaseAdapter {
	return JobcaseAdapter{
		core: core,
		dao:  dao,
	}
}

func (adapter JobcaseAdapter) SaveJob(job job.Job) bool {
	if !adapter.dao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v save error", job.ID)
		return false
	}
	return true
}

func (adapter JobcaseAdapter) DeleteJob(job job.Job) bool {
	if !adapter.dao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v delete error", job.ID)
		return false
	}
	return true
}

func (adapter JobcaseAdapter) FindJobByID(id string) job.Job {
	return adapter.dao.FindJobByID(id)
}

func (adapter JobcaseAdapter) UpdateJob(job job.Job) bool {
	if !adapter.dao.UpdateJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v update error", job.ID)
		return false
	}
	return true
}

func (adapter JobcaseAdapter) ApplyJob() {
	// 收集用户提交信息

	// 验证信息

	// 保存信息至数据库，搭建关联表

}
