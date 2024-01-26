package user

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	jobApp "github.com/Renewdxin/selfMade/internal/ports/app/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type AdminAppAdapter struct {
	job    jobApp.JobsCasePorts
	jobDao database.JobDaoPorts
	user   database.UserDaoPorts
}

func NewAdapter(job jobApp.JobsCasePorts, jobDao database.JobDaoPorts, user database.UserDaoPorts) AdminAppAdapter {
	return AdminAppAdapter{
		job:    job,
		jobDao: jobDao,
		user:   user,
	}
}

func (adapter AdminAppAdapter) ShowJobsDetails(id string) job.Job {
	return adapter.job.FindJobByID(id)
}

func (adapter AdminAppAdapter) ShowAllJobs() []job.Job {
	return adapter.jobDao.ShowAllJobs()
}

func (adapter AdminAppAdapter) ShowJobsApply() []user.User {
	return []user.User{}
}

func (adapter AdminAppAdapter) ApproveJobs(u user.User) bool {
	if !adapter.user.ChangeUserStatus(u.ID, 1-u.State) {
		logger.Logger.Logf(logger.ErrorLevel, "failed to change user's status, userID: %v", u.ID)
		return false
	}
	logger.Logger.Logf(logger.InfoLevel, "change user's status, userID: %v", u.ID)
	return true
}
