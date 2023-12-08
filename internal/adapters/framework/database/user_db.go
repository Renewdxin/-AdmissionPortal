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
	if err := userDao.db.Create(newUser).Error; err != nil {
		log.Fatalf("Failed to save : %v", newUser)
		return nil, err
	}

	return newUser, nil
}

func (userDao *UserDao) DeleteUser(id string) error {

}

func (userDao *UserDao) UpdateUser(id string) (*user.User, error) {

}

func (userDao *UserDao) FindUserByID(id string) (*user.User, error) {

}

func (userDao *UserDao) FindUserByEmail(email string) (*user.User, error) {

}
