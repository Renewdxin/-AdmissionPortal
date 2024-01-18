package database

import (
	"database/sql"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AdminDaoAdapter struct {
	db *gorm.DB
}

func NewAdminDaoAdapter(driveName, dataSourceName string) AdminDaoAdapter {
	sqlDB, err := sql.Open(driveName, dataSourceName)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "init error: %v", err)
		return AdminDaoAdapter{}
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "init error: %v", err)
		return AdminDaoAdapter{}
	}
	return AdminDaoAdapter{gormDB}
}
