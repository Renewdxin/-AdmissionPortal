package database

import (
	"database/sql"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AdminDaoAdapter struct {
	db        *gorm.DB
	adminCore user.AdminCorePorts
}

func NewAdminDaoAdapter(driveName, dataSourceName string, adminCore user.AdminCorePorts) AdminDaoAdapter {
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
	return AdminDaoAdapter{gormDB, adminCore}
}

func (dao AdminDaoAdapter) SaveAdminAccount(account auth.Account) bool {
	if err := dao.db.Table(dao.adminCore.TableName()).Create(&account).Error; err != nil {
		return false
	}
	return true
}

func (dao AdminDaoAdapter) UpdateAccount(account auth.Account) bool {
	result := dao.db.Table(dao.adminCore.TableName()).Model(&account).Updates(map[string]interface{}{
		"id":       account.ID,
		"password": account.Password,
	})
	if result.Error != nil {
		return false
	}

	return true
}

func (dao AdminDaoAdapter) DeleteAccount(id string) bool {
	if err := dao.db.Table(dao.adminCore.TableName()).Delete(&auth.Account{}, id); err != nil {

		return false
	}
	return true
}

func (dao AdminDaoAdapter) FindAccountByID(id string) bool {
	var account auth.Account
	if err := dao.db.Table(dao.adminCore.TableName()).Where("id = ?", id).First(&account).Error; err != nil {
		return false
	}
	return true
}
