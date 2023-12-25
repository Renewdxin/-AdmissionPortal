package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"time"
)

type UserCore struct{}

func NewUserService() *UserCore {
	return &UserCore{}
}

func (u UserCore) CreateUser(name string, gender string, email string, phone string, birth string) (user.User, error) {
	// validate the u
	newUser := user.User{
		Name:        name,
		CreatedAt:   time.Now(),
		State:       0,
		Birth:       birth,
		Gender:      gender,
		Email:       email,
		PhoneNumber: phone,
	}
	return newUser, nil
}

func (u UserCore) TableName() string {
	return "userinfo"
}
