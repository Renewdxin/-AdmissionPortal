package job

import "github.com/Renewdxin/selfMade/internal/ports/core/job"

type JobsAdapter struct{}

func NewAdapter() JobsAdapter {
	return JobsAdapter{}
}

func (adapter JobsAdapter) CreateJob(id, name string) job.Job {
	return job.Job{ID: id, Name: name}
}

func (adapter JobsAdapter) TableName() string {
	return "job"
}

func (adapter JobsAdapter) ShowJob(job job.Job) (id, name string) {
	return job.ID, job.Name
}
