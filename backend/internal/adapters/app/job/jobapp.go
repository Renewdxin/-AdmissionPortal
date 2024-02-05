package job

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type JobsApplicationAdapter struct {
	core job.JobsCorePort
	dao  database.JobDBPort
}

func NewJobCaseAdapter(core job.JobsCorePort, dao database.JobDBPort) JobsApplicationAdapter {
	return JobsApplicationAdapter{
		core: core,
		dao:  dao,
	}
}

func (adapter JobsApplicationAdapter) SaveJob(job job.Job) bool {
	if !adapter.dao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v save error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) DeleteJob(job job.Job) bool {
	if !adapter.dao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v delete error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) FindJobByID(id string) job.Job {
	return adapter.dao.FindJobByID(id)
}

func (adapter JobsApplicationAdapter) UpdateJob(job job.Job) bool {
	if !adapter.dao.UpdateJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v update error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) ApplyJob() {
	// 收集用户提交信息

	// 验证信息

	// 保存信息至数据库，搭建关联表

}

func (adapter JobsApplicationAdapter) ShowAllJobs() []job.Job {
	return adapter.dao.ShowAllJobs()
}
