package database

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserDao interface {
	SaveUser(name string, gender string, email string, phone string) (*user.User, error)
	DeleteUser(id string) error
	UpdateUser(id string) (*user.User, error)
	FindUserByID(id string) (*user.User, error)
	FindUserByEmail(email string) (*user.User, error)
}
