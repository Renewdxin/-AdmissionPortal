package database

import (
	"database/sql"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/job"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JobsDaoAdapter struct {
	db   *gorm.DB
	core job.JobsCorePorts
}

func NewJobsDaoAdapter(driveName, dataSourceName string, core job.JobsCorePorts) JobsDaoAdapter {
	sqlDB, err := sql.Open(driveName, dataSourceName)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "init error: %v", err)
		return JobsDaoAdapter{}
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "init error: %v", err)
		return JobsDaoAdapter{}
	}
	return JobsDaoAdapter{gormDB, core}
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

func (dao JobsDaoAdapter) ShowAllJobs() []job.Job {
	var jobSet []job.Job
	if err := dao.db.Find(&jobSet).Error; err != nil {
		logger.Logger.Log(logger.WarnLevel, "Failed to show all jobs")
		return []job.Job{}
	}
	return jobSet
}
