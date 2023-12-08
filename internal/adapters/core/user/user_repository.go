package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type UserRepoAdapter struct {
	userDao database.UserDao
}

func NewUserRepoAdapter() *UserRepoAdapter {
	return &UserRepoAdapter{}
}

func (userRepo *UserRepoAdapter) CreateUser(name string, gender string, email string, phone string) (*user.User, error) {
	newUser, _ := userRepo.userDao.SaveUser(name, gender, email, phone)
	return newUser, nil
}

func (userRepo *UserRepoAdapter) FindByID(id string) (*user.User, error) {
	newUser, _ := userRepo.userDao.FindUserByID(id)
	return newUser, nil
}

func (userRepo *UserRepoAdapter) FindByEmail(email string) (*user.User, error) {
	newUser, _ := userRepo.userDao.FindUserByEmail(email)
	return newUser, nil
}

func (userRepo *UserRepoAdapter) Update(id string) (*user.User, error) {
	newUser, _ := userRepo.userDao.UpdateUser(id)
	return newUser, nil
}

func (userRepo *UserRepoAdapter) Delete(id string) error {
	err := userRepo.userDao.DeleteUser(id)
	return err
}
