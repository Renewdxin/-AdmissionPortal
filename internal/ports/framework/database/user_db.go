package database

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserDaoPorts interface {
	SaveUser(name string, gender string, email string, phone string) (*user.User, error)
	DeleteUser(id string) error
	UpdateUser(*user.User) (*user.User, error)
	FindUserByID(id string) (*user.User, error)
	FindUserByEmail(email string) (*user.User, error)
	IfExist(email string) bool
}
