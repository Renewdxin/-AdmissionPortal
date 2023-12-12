package auth

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type AuthcaseAdapter struct {
	core auth.AccountCorePorts
	dao  database.AuthDaoPorts
}

func NewAuthcaseAdapter(core auth.AccountCorePorts, dao database.AuthDaoPorts) *AuthcaseAdapter {
	return &AuthcaseAdapter{
		core: core,
		dao:  dao,
	}
}

func (auth AuthcaseAdapter) Register(id, password string) error {
	if auth.dao.SaveAccount(id, password) {
		return errors.New("failed to register, plz try again")
	}
	return nil
}
func (auth AuthcaseAdapter) ForgetPassword(id, phone, password string) error {
	if auth.dao.UpdateAccount(id, password) {
		return errors.New("failed to change, plz try again")
	}
	return nil

}

func (auth AuthcaseAdapter) ChangePassword(id, oldPassword, newPassword string) error {
	if auth.dao.UpdateAccount(id, newPassword) {
		return errors.New("failed to change, plz try again")
	}
	return nil

}

func (auth AuthcaseAdapter) LogIn(id, password string) error {
	if auth.dao.UpdateAccount(id, password) {
		return errors.New("failed to change, plz try again")
	}
	return nil

}
