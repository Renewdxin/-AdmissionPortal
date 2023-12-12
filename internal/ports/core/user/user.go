package user

import (
	"time"
)

type User struct {
	ID          string `json:"id" gorm:"primarykey;column:id" `
	Name        string `json:"name" gorm:"column:name"`
	CreatedAt   time.Time
	State       int    `json:"state" validate:"oneof=0 1"`
	Gender      string `json:"gender" gorm:"column:gender"`
	Birth       string `json:"birth"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone"`
	//Account     auth.Account
}

func (user User) TableName() string {
	return "userinfo"
}

type UserPorts interface {
	CreateUser(name string, gender string, email string, phone string) (User, error)
	TableName() string
}
