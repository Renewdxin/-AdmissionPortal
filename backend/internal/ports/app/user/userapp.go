package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UsrApplicationPort interface {
	RegisterUser(name, gender, email, phone, birth string) error
	GetUserProfile(id string) (user.User, error)
	DeleteUser(id string) error
	UpdateUser(user user.User) error
	IfExist(email string) bool
	UserValidateBeforeRegister(name string, gender string, email string, phone string, birth string) bool
}
