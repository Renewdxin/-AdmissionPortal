package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
)

type UserCore struct{}

func NewUserService() *UserCore {
	return &UserCore{}
}

func (u UserCore) CreateUser(name string, gender string, email string, phone string) (user.User, error) {
	// validate the u
	newUser := user.User{
		Name:        name,
		Email:       email,
		PhoneNumber: phone,
		Gender:      gender,
	}
	return newUser, nil
}

func (u UserCore) TableName() string {
	return "userinfo"
}
