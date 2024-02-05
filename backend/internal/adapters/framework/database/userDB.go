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
func (userDao UsrDaoAdapter) IfExist(email string) bool {
	var u user.User
	if err := userDao.db.Table(userDao.user.TableName()).Where("email = ?", email).First(&u).Error; err != nil {
		return false
	}
	return true
}

func (userDao UsrDaoAdapter) SaveUser(user user.User) error {
	if err := userDao.db.Table(userDao.user.TableName()).Select("name", "gender", "email", "phone").Create(&user).Error; err != nil {
		log.Printf("Failed to save user %v: %v", user.Name, err)
		return err
	}
	return nil
}

func (userDao UsrDaoAdapter) DeleteUser(id string) error {
	result := userDao.db.Table(userDao.user.TableName()).Delete(&user.User{}, id)
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

func (userDao UsrDaoAdapter) UpdateUser(user user.User) error {
	result := userDao.db.Table(userDao.user.TableName()).Model(&user).Updates(map[string]interface{}{
		"name":   user.Name,
		"gender": user.Gender,
		"email":  user.Email,
		"phone":  user.PhoneNumber,
	})
	if result.Error != nil {
		log.Printf("Failed to update user: %v", result.Error)
		return result.Error
	}
	return nil
}

func (userDao UsrDaoAdapter) FindUserByID(id string) (user.User, error) {
	var newUser user.User
	if err := userDao.db.Table(userDao.user.TableName()).Where("id = ?", id).First(&newUser).Error; err != nil {
		log.Printf("Failed to find user by ID %v: %v", id, err)
		return user.User{}, err
	}
	return newUser, nil
}

func (userDao UsrDaoAdapter) FindUserByEmail(email string) (user.User, error) {
	var newUser user.User
	if err := userDao.db.Table(userDao.user.TableName()).Where("email = ?", email).First(&newUser).Error; err != nil {
		log.Printf("Failed to find user by email %v: %v", email, err)
		return user.User{}, err
	}
	return newUser, nil
}

func (userDao UsrDaoAdapter) ChangeUserStatus(id string, state int) bool {
	update := user.User{State: state}
	if err := userDao.db.Table(userDao.user.TableName()).Where("id = ?", id).Updates(update).Error; err != nil {
		log.Printf("Failed to find user by ID %v: %v", id, err)
		return false
	}
	return true
}
