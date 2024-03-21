package user

import (
	"time"
)

type User struct {
	ID          string    `json:"id" gorm:"primaryKey;column:id" `
	Academy     string    `json:"academy" gorm:"column:academy"`
	Major       string    `json:"major" gorm:"column:major"`
	Name        string    `json:"name" gorm:"column:name"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	State       int       `json:"state" validate:"oneof=0 1" gorm:"default:0"`
	Email       string    `json:"email" gorm:"column:email"`
	PhoneNumber string    `json:"phoneNumber" gorm:"column:phone"`
	ApplyID     int       `json:"apply_id" gorm:"column:apply_id"`
	//Account     auth.Account `json:"account"`
}

func (user User) TableName() string {
	return "userinfo"
}

type UsrCorePort interface {
	CreateUser(name string, gender string, email string, phone string, birth string) (User, error)
	TableName() string
}
