package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
)

type AdminApplicationPort interface {
	ShowJobsDetails(id string) job.Job
	ShowAllJobs() []job.Job
	ShowJobsApply() ([]user.User, error)
	ApproveJobs(u user.User) bool
	ShowJobApplyByJobID(jobID string) ([]user.User, error)
	ShowAllUnhandledApplications() ([]user.User, error)
	ShowUnhandledApplicationsByJobID(jobID string) ([]user.User, error)
}
