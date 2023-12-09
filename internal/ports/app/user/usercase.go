package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserCasePorts interface {
	RegisterUser(name, gender, email, phone string) error
	GetUserProfile(id uint) (*user.User, error)
	IfExist(email string) bool
}
