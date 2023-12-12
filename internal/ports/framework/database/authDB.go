package database

type AuthDaoPorts interface {
	SaveAccount(id, password string) bool
	UpdateAccount(id, password string) bool
	DeleteAccount(id string) bool
	FindAccountByID(id string) bool
}
