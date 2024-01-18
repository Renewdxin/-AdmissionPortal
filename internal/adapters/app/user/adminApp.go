package user

import (
	jobApp "github.com/Renewdxin/selfMade/internal/ports/app/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type AdminAppAdapter struct {
	job    jobApp.JobsCasePorts
	jobDao database.JobDaoPorts
}

func NewAdapter() AdminAppAdapter {
	return AdminAppAdapter{}
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

func (adapter AdminAppAdapter) ApproveJobs(u user.User) {

}
