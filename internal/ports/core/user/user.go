package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/account"
	"time"
)

type User struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `json:"name"`
	CreatedAt   time.Time
	State       int    `json:"state" gorm:"type:tinyint" validate:"oneof=0 1"`
	Gender      string `json:"gender"`
	Birth       string `json:"birth"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Account     account.Account
}

func (user *User) TableName() string {
	return "user"
}
