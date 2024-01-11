package job

import "github.com/Renewdxin/selfMade/internal/ports/core/job"

type jobCasePorts interface {
	SaveJob(job job.Job) bool
	DeleteJob(job job.Job) bool
	FindJobByID(id string) job.Job
	UpdateJob(job job.Job) bool
	ApplyJob()
}
