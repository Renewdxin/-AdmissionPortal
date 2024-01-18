package database

import "github.com/Renewdxin/selfMade/internal/ports/core/job"

type JobDaoPorts interface {
	SaveJob(job job.Job) bool
	DeleteJob(id string) bool
	UpdateJob(job job.Job) bool
	FindJobByID(id string) job.Job
	ShowAllJobs() []job.Job
}
