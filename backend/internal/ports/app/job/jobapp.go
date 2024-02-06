package job

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
)

type JobsApplicationPort interface {
	SaveJob(job job.Job) bool
	DeleteJob(job job.Job) bool
	FindJobByID(id string) job.Job
	UpdateJob(job job.Job) bool
	ApplyJob(usr user.User, jobID int, userID string) bool
	ShowAllJobs() []job.Job
}
