package user

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserRepoAdapter struct {
}

func NewUserRepoAdapter() *UserRepoAdapter {
	return &UserRepoAdapter{}
}

func (userRepo *UserRepoAdapter) CreateUser(name string, gender string, email string, phone string) (*user.User, error) {

	return nil
}

func (userRepo *UserRepoAdapter) FindByID(id string) (*user.User, error) {
	return nil
}

func (userRepo *UserRepoAdapter) Update(id string) (*user.User, error) {
	return nil
}

func (userRepo *UserRepoAdapter) Delete(id string) error {
	return nil
}

func (userRepo *UserRepoAdapter) FindByEmail(id string) (*user.User, error) {
	return nil
}
