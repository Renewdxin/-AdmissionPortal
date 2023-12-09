package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserCasePorts interface {
	RegisterUser(name, gender, email, phone string) error
	GetUserProfile(id int) (*user.User, error)
	DeleteUser(id string) error
	UpdateUser(user user.User) error
	IfExist(email string) bool
}
