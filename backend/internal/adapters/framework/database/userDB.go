package database

import (
	"database/sql"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type UsrDaoAdapter struct {
	db   *gorm.DB
	user user.UsrCorePort
}

func NewUsrDaoAdapter(driveName, dataSourceName string, user user.UsrCorePort) (*UsrDaoAdapter, error) {
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
	return &UsrDaoAdapter{db: gormDB, user: user}, nil
}

// IfExist
// true exist
func (dao UsrDaoAdapter) IfExist(email string) bool {
	var u user.User
	if err := dao.db.Table(dao.user.TableName()).Where("email = ?", email).First(&u).Error; err != nil {
		return false
	}
	return true
}

func (dao UsrDaoAdapter) SaveUser(user user.User) error {
	if err := dao.db.Table(dao.user.TableName()).Create(&user).Error; err != nil {
		log.Printf("Failed to save user %v: %v", user.Name, err)
		return err
	}
	return nil
}

func (dao UsrDaoAdapter) DeleteUser(id string) error {
	result := dao.db.Table(dao.user.TableName()).Delete(&user.User{}, id)
	if result.Error != nil {
		log.Printf("Failed to delete user with ID %v: %v", id, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Printf("User with ID %v not found", id)
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (dao UsrDaoAdapter) UpdateUser(user user.User) error {
	if err := dao.db.Table(dao.user.TableName()).Model(&user).Updates(map[string]interface{}{
		"apply_id": user.ApplyID}).Error; err != nil {
		return err
	}
	return nil
}

func (dao UsrDaoAdapter) FindUserByID(id string) (user.User, error) {
	var newUser user.User
	if err := dao.db.Table(dao.user.TableName()).Where("id = ?", id).First(&newUser).Error; err != nil {
		return user.User{}, err
	}
	return newUser, nil
}

func (dao UsrDaoAdapter) FindUserByEmail(email string) (user.User, error) {
	var newUser user.User
	if err := dao.db.Table(dao.user.TableName()).Where("email = ?", email).First(&newUser).Error; err != nil {
		return user.User{}, err
	}
	return newUser, nil
}

func (dao UsrDaoAdapter) ChangeUserStatus(id string, state int) bool {
	update := user.User{State: state}
	if err := dao.db.Table(dao.user.TableName()).Where("id = ?", id).Updates(update).Error; err != nil {
		return false
	}
	return true
}

// ShowAllUsers 返回数据库中所有用户的切片。
func (dao UsrDaoAdapter) ShowAllUsers() ([]user.User, error) {
	var users []user.User
	if err := dao.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ShowJobApplyByJobID 根据岗位ID返回对应岗位的申请者的切片。
func (dao UsrDaoAdapter) ShowJobApplyByJobID(jobID string) ([]user.User, error) {
	var users []user.User
	if err := dao.db.Where("job_id = ?", jobID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ShowAllUnhandledApplications 返回所有未处理的申请者的切片。
func (dao UsrDaoAdapter) ShowAllUnhandledApplications() ([]user.User, error) {
	var users []user.User
	if err := dao.db.Where("state = ?", "UNHANDLED").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ShowUnhandledApplicationsByJobID 根据岗位ID返回未处理的申请者的切片。
func (dao UsrDaoAdapter) ShowUnhandledApplicationsByJobID(jobID string) ([]user.User, error) {
	var users []user.User
	if err := dao.db.Where("job_id = ? AND state = ?", jobID, "UNHANDLED").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
