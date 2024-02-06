package job

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type JobsApplicationAdapter struct {
	core   job.JobsCorePort
	jobDao database.JobDBPort
	usrDao database.UserDBPort
}

func NewJobCaseAdapter(core job.JobsCorePort, dao database.JobDBPort, usrDao database.UserDBPort) JobsApplicationAdapter {
	return JobsApplicationAdapter{
		core:   core,
		jobDao: dao,
		usrDao: usrDao,
	}
}

func (adapter JobsApplicationAdapter) SaveJob(job job.Job) bool {
	if !adapter.jobDao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v save error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) DeleteJob(job job.Job) bool {
	if !adapter.jobDao.SaveJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v delete error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) FindJobByID(id string) job.Job {
	return adapter.jobDao.FindJobByID(id)
}

func (adapter JobsApplicationAdapter) UpdateJob(job job.Job) bool {
	if !adapter.jobDao.UpdateJob(job) {
		logger.Logger.Logf(logger.InfoLevel, "job: %v update error", job.ID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) ApplyJob(usr user.User, jobID int, userID string) bool {
	if jobID < 0 || jobID > 4 {
		logger.Logger.Logf(logger.InfoLevel, "failed to apply job : %v user : %v", jobID, userID)
		return false
	}
	// 查找用户是否已注册
	var err error
	_, err = adapter.usrDao.FindUserByID(userID)
	// 第一种情况 ： 用户未注册过
	if err != nil {
		usr.ID = userID
		usr.Account.ID = userID
		usr.ApplyID = jobID
		err := adapter.usrDao.SaveUser(usr)
		if err != nil {
			logger.Logger.Logf(logger.InfoLevel, "failed to apply job : %v user : %v", jobID, userID)
			return false
		}
		logger.Logger.Logf(logger.InfoLevel, "new user : %v apply the job : %v ", userID, jobID)
	}
	// 第二种情况 ： 已注册
	// 已经注册过方向:予以驳回
	if usr.ApplyID != 10 {
		logger.Logger.Logf(logger.InfoLevel, "already apply the job : %v user : %v", jobID, userID)
		return false
	}
	usr.ApplyID = jobID
	err = adapter.usrDao.SaveUser(usr)
	if err != nil {
		logger.Logger.Logf(logger.InfoLevel, "failed to apply job : %v user : %v", jobID, userID)
		return false
	}
	return true
}

func (adapter JobsApplicationAdapter) ShowAllJobs() []job.Job {
	return adapter.jobDao.ShowAllJobs()
}
