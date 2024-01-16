package database

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
)

type adminDBPorts interface {
	ShowJobsDetails(id string) job.Job
	ShowAllJobs() []job.Job
	ShowJobsApply() []user.User
	ApproveJobs(id string) string
}
