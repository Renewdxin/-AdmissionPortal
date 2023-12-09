package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type userAPIPorts interface {
	RegisterUser(name, gender, email, phone string) (*user.User, error)
	GetUserProfile(id uint) (*user.User, error)
	IfExist(email string) bool
}
