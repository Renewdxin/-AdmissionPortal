package auth

type Account struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
}

type AccountCorePorts interface {
	CreateAccount(id, password string) (Account, error)
	TableName() string
}
