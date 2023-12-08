package user

type UserRepository interface {
	CreateUser(name string, gender string, email string, phone string) (*User, error)
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
}
