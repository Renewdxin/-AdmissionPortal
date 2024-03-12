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

func (api AdminApplicationAdapter) ShowJobsDetails(id string) job.Job {
	return api.job.FindJobByID(id)
}

func (api AdminApplicationAdapter) ShowAllJobs() []job.Job {
	return api.jobDao.ShowAllJobs()
}

func (api AdminApplicationAdapter) ShowJobsApply() ([]user.User, error) {
	return api.user.ShowAllUsers()
}

func (api AdminApplicationAdapter) ApproveJobs(u user.User) bool {
	if !api.user.ChangeUserStatus(u.ID, 1-u.State) {
		logger.Logger.Logf(logger.ErrorLevel, "failed to change user's status, userID: %v", u.ID)
		return false
	}
	logger.Logger.Logf(logger.InfoLevel, "user's status change , userID: %v", u.ID)
	return true
}

func (api AdminApplicationAdapter) ShowJobApplyByJobID(jobID string) ([]user.User, error) {
	return api.user.ShowJobApplyByJobID(jobID)
}

func (api AdminApplicationAdapter) ShowAllUnhandledApplications() ([]user.User, error) {
	return api.user.ShowAllUnhandledApplications()
}

func (api AdminApplicationAdapter) ShowUnhandledApplicationsByJobID(jobID string) ([]user.User, error) {
	return api.user.ShowUnhandledApplicationsByJobID(jobID)
}
