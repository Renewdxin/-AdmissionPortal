package database

type AuthDaoAdapter struct{}

func NewauthDaoAdapter() *AuthDaoAdapter {
	return &AuthDaoAdapter{}
}

func (dao AuthDaoAdapter) SaveAccount(id, password string) bool {
	return true
}

func (dao AuthDaoAdapter) UpdateAccount(id, password string) bool {
	return true
}

func (dao AuthDaoAdapter) DeleteAccount(id string) bool {
	return true
}

func (dao AuthDaoAdapter) FindAccountByID(id string) bool {
	return true
}
