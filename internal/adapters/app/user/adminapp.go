package user

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	jobApp "github.com/Renewdxin/selfMade/internal/ports/app/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type AdminApplicationAdapter struct {
	job    jobApp.JobsApplicationPort
	jobDao database.JobDBPort
	user   database.UserDBPort
}

func NewAdminAppAdapter(job jobApp.JobsApplicationPort, jobDao database.JobDBPort, user database.UserDBPort) AdminApplicationAdapter {
	return AdminApplicationAdapter{
		job:    job,
		jobDao: jobDao,
		user:   user,
	}
}

func (adapter AdminApplicationAdapter) ShowJobsDetails(id string) job.Job {
	return adapter.job.FindJobByID(id)
}

func (adapter AdminApplicationAdapter) ShowAllJobs() []job.Job {
	return adapter.jobDao.ShowAllJobs()
}

func (adapter AdminApplicationAdapter) ShowJobsApply() []user.User {
	return []user.User{}
}

func (adapter AdminApplicationAdapter) ApproveJobs(u user.User) bool {
	if !adapter.user.ChangeUserStatus(u.ID, 1-u.State) {
		logger.Logger.Logf(logger.ErrorLevel, "failed to change user's status, userID: %v", u.ID)
		return false
	}
	logger.Logger.Logf(logger.InfoLevel, "change user's status, userID: %v", u.ID)
	return true
}
