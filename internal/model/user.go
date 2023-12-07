package model

import (
	"time"
)

type User struct {
	ID          uint `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time
	State       int    `json:"state" gorm:"type:tinyint" validate:"oneof=0 1"`
	Gender      string `json:"gender"`
	Birth       string `json:"birth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Account     Account
}

func (employee *User) TableName() string {
	return "employee"
}
