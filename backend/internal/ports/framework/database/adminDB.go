package database

import "github.com/Renewdxin/selfMade/internal/ports/core/auth"

type AdminDBPort interface {
	SaveAdminAccount(account auth.Account) bool
	FindAccountByID(id string) bool
	UpdateAccount(account auth.Account) bool
	DeleteAccount(id string) bool
}
