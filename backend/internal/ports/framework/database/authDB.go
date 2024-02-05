package database

import "github.com/Renewdxin/selfMade/internal/ports/core/auth"

type AuthDBPort interface {
	EmailIfExist(email string) bool
	SaveUserAccount(account auth.Account) bool
	UpdateAccount(account auth.Account) bool
	DeleteAccount(id string) bool
	FindAccountByID(id string) bool
	FindPasswordByID(id string) string
}
