package ports

type UserRepository interface {
	Create(id string) error
	Find(id string) error
	FindByEmail(email string) error
	Update(id string) error
	Delete(id string) error
}
