package auth

import "github.com/Renewdxin/selfMade/internal/ports/core/auth"

type AuthorizeCoreAdapter struct{}

func NewAuthorizeCoreAdapter() AuthorizeCoreAdapter {
	return AuthorizeCoreAdapter{}
}

func (accountAdapter AuthorizeCoreAdapter) CreateAccount(id, password string) (auth.Account, error) {
	return auth.Account{ID: id, Password: password}, nil
}

func (accountAdapter AuthorizeCoreAdapter) TableName() string {
	return "user_account"
}
