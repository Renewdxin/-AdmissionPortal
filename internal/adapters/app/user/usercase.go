package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserUseCase interface {
	RegisterUser(username, email, password string) (*user.User, error)
	GetUserProfile(userID uint) (*user.User, error)
	// 其他用户用例方法
}
