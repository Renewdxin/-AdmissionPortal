package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"log"
)

type UserService struct {
	repo user.RepositoryPorts
}

func NewUserService(repositoryPorts user.RepositoryPorts) *UserService {
	return &UserService{
		repo: repositoryPorts,
	}
}

func (user *UserService) CreateUser(name string, gender string, email string, phone string) (*user.User, error) {
	if len(name) > 50 || len(name) < 2 {
		log.Fatalln("INVALID NAME")
	}

	newUser, err := user.repo.CreateUser(name, gender, email, phone)
	if err != nil {
		log.Fatalf("Failed to create the User: %v, plz try again", name)
	}
	return newUser, nil
}

func (user *UserService) FindByID(id string) (*user.User, error) {
	newUser, err := user.repo.FindByID(id)
	if err != nil {
		log.Fatalf("Failed to find the User with the id: %v, plz try again", id)
	}
	return newUser, nil
}

func (user *UserService) FindByEmail(email string) (*user.User, error) {
	newUser, err := user.repo.FindByEmail(email)
	if err != nil {
		log.Fatalf("Failed to find the User with the email: %v, plz try again", email)
	}
	return newUser, nil
}

func (user *UserService) Update(newUser *user.User) (*user.User, error) {
	data, err := user.repo.Update(newUser)
	if err != nil {
		log.Fatalf("Failed to update the User: %v, plz try again", newUser.Name)
	}
	return data, nil
}

func (user *UserService) Delete(id string) error {
	err := user.repo.Delete(id)
	if err != nil {
		log.Fatalf("Failed to update the User with the id: %v, plz try again", id)
		return err
	}
	return nil
}
