package database

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"gorm.io/gorm"
)

type JobsDaoAdapter struct {
	db   *gorm.DB
	core job.JobsCorePorts
}

func NewAdapter() JobsDaoAdapter {
	return JobsDaoAdapter{}
}

func (dao JobsDaoAdapter) SaveJob(job job.Job) bool {
	if err := dao.db.Table(dao.core.TableName()).Create(&job).Error; err != nil {
		return false
	}
	return true
}

func (dao JobsDaoAdapter) DeleteJob(id string) bool {
	if err := dao.db.Table(dao.core.TableName()).Delete(&job.Job{}, id); err != nil {
		return false
	}
	return true
}

func (dao JobsDaoAdapter) UpdateJob(job job.Job) bool {
	if err := dao.db.Table(dao.core.TableName()).Create(&job).Error; err != nil {
		return false
	}
	return true
}

func (dao JobsDaoAdapter) FindJobByID(id string) job.Job {
	var jobInfo job.Job
	if err := dao.db.Table(dao.core.TableName()).Where("id = ?", id).First(&jobInfo).Error; err != nil {
		return jobInfo
	}
	return job.Job{}
}
