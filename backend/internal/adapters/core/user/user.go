package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"time"
)

type UsrCoreAdapter struct{}

func NewUsrCoreAdapter() *UsrCoreAdapter {
	return &UsrCoreAdapter{}
}

func (u UsrCoreAdapter) CreateUser(name string, gender string, email string, phone string, birth string) (user.User, error) {
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

func (u UsrCoreAdapter) TableName() string {
	return "userInfo"
}
