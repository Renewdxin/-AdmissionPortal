package model

type Account struct {
	ID       string `gorm:"primarykey" json:"id"`
	Password string `json:"password"`
}

func (account *Account) AccountTableName() string {
	return "accounts"
}
