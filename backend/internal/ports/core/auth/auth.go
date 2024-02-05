package auth

type Account struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
}

type AuthorizeCorePort interface {
	CreateAccount(id, password string) (Account, error)
	TableName() string
}
