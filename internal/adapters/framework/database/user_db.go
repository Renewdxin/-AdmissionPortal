package database

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	db gorm.DB
}

func NewUserDao(db gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

// IfExist
// true exist
func (userDao *UserDao) IfExist(email string) bool {
	var u user.User
	if err := userDao.db.Where("email = ?", email).First(&u).Error; err != nil {
		return false
	}
	return true
}

func (userDao *UserDao) SaveUser(user user.User) error {
	if err := userDao.db.Select("name", "gender", "email", "phone").Create(&user).Error; err != nil {
		log.Fatalf("Failed to save : %v", user.Name)
		return err
	}
	return nil
}

func (userDao *UserDao) DeleteUser(id string) error {
	if err := userDao.db.Delete(&user.User{}, id).Error; err != nil {
		log.Fatalln("Failed to delete")
		return err
	}
	return nil
}

func (userDao *UserDao) UpdateUser(user user.User) error {
	if err := userDao.db.Model(user).Updates(user).Error; err != nil {
		log.Fatalln("Failed to update")
		return err
	}
	return nil
}

func (userDao *UserDao) FindUserByID(id string) (*user.User, error) {
	var newUser user.User
	if err := userDao.db.First(&newUser, id).Error; err != nil {
		log.Fatalf("failed to find user by id: %v", id)
		return nil, err
	}
	return &newUser, nil
}

func (userDao *UserDao) FindUserByEmail(email string) (*user.User, error) {
	var newUser user.User
	if err := userDao.db.Where("email = ?", email).First(&newUser).Error; err != nil {
		log.Fatalf("failed to find user by email: %v", email)
		return nil, err
	}
	return &newUser, nil
}
