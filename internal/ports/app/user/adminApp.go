package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
)

type AdminApplicationPorts interface {
	ShowJobsDetails(id string) job.Job
	ShowAllJobs() []job.Job
	ShowJobsApply() []user.User
	ApproveJobs(u user.User) bool
}
