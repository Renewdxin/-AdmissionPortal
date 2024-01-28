package job

import "github.com/Renewdxin/selfMade/internal/ports/core/job"

type JobsCoreAdapter struct{}

func NewJobsCoreAdapter() JobsCoreAdapter {
	return JobsCoreAdapter{}
}

func (adapter JobsCoreAdapter) CreateJob(id, name string) job.Job {
	return job.Job{ID: id, Name: name}
}

func (adapter JobsCoreAdapter) TableName() string {
	return "job"
}

func (adapter JobsCoreAdapter) ShowJob(job job.Job) (id, name string) {
	return job.ID, job.Name
}
