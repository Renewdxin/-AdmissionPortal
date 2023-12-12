package auth

type Account struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type AccountCorePorts interface {
	CreateAccount(id, password string) (*Account, error)
	TableName() string
}
