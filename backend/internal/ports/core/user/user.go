package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"time"
)

type User struct {
	ID          string       `json:"id" gorm:"primaryKey;column:id" `
	Name        string       `json:"name" gorm:"column:name"`
	CreatedAt   time.Time    `gorm:"column:created_at"`
	State       int          `json:"state" validate:"oneof=0 1" gorm:"default:0"`
	Gender      string       `json:"gender" gorm:"column:gender"`
	Birth       string       `json:"birth"`
	Email       string       `json:"email" gorm:"column:email"`
	PhoneNumber string       `json:"phoneNumber" gorm:"column:phone"`
	ApplyID     int          `json:"applyID"`
	Account     auth.Account `json:"account"`
}

func (user User) TableName() string {
	return "userinfo"
}

type UsrCorePort interface {
	CreateUser(name string, gender string, email string, phone string, birth string) (User, error)
	TableName() string
}
