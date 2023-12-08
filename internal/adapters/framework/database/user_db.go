package database

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (userDao *UserDao) SaveUser(name string, gender string, email string, phone string) (*user.User, error) {
	newUser := &user.User{
		Name:        name,
		Gender:      gender,
		Email:       email,
		PhoneNumber: phone,
	}
	if err := userDao.db.Select("name", "gender", "email", "phone").Create(newUser).Error; err != nil {
		log.Fatalf("Failed to save : %v", newUser)
		return nil, err
	}
	return newUser, nil
}

func (userDao *UserDao) DeleteUser(id string) error {
	if err := userDao.db.Delete(&user.User{}, id).Error; err != nil {
		log.Fatalln("Failed to delete")
		return err
	}
	return nil
}

func (userDao *UserDao) UpdateUser(user *user.User) (*user.User, error) {
	if err := userDao.db.Model(user).Updates(user).Error; err != nil {
		log.Fatalln("Failed to update")
		return nil, err
	}
	return user, nil
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
