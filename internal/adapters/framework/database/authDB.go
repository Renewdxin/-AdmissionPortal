package database

import (
	"database/sql"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthDaoAdapter struct {
	db      *gorm.DB
	account auth.AuthorizeCorePort
}

func NewAuthDaoAdapter(driveName, dataSourceName string, account auth.AuthorizeCorePort) (*AuthDaoAdapter, error) {
	sqlDB, err := sql.Open(driveName, dataSourceName)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &AuthDaoAdapter{db: gormDB, account: account}, nil
}

func (dao AuthDaoAdapter) EmailIfExist(email string) bool {
	var account auth.Account
	if err := dao.db.Table(dao.account.TableName()).Where("email = ?", email).Find(&account).Error; err != nil {
		return false
	}
	return true
}

func (dao AuthDaoAdapter) SaveUserAccount(account auth.Account) bool {
	if err := dao.db.Table(dao.account.TableName()).Create(&account).Error; err != nil {
		return false
	}
	return true
}

func (dao AuthDaoAdapter) UpdateAccount(account auth.Account) bool {
	result := dao.db.Table(dao.account.TableName()).Model(&account).Updates(map[string]interface{}{
		"id":       account.ID,
		"password": account.Password,
	})
	if result.Error != nil {
		return false
	}

	return true
}

func (dao AuthDaoAdapter) DeleteAccount(id string) bool {
	if err := dao.db.Table(dao.account.TableName()).Delete(&auth.Account{}, id); err != nil {

		return false
	}
	return true
}

func (dao AuthDaoAdapter) FindAccountByID(id string) bool {
	var account auth.Account
	if err := dao.db.Table(dao.account.TableName()).Where("id = ?", id).First(&account).Error; err != nil {
		return false
	}
	return true
}

func (dao AuthDaoAdapter) FindPasswordByID(id string) string {
	var account auth.Account
	if err := dao.db.Table(dao.account.TableName()).Where("id = ?", id).First(&account).Error; err != nil {
		return ""
	}
	return account.Password
}
