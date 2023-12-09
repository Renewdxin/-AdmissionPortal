package user

import (
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type UserRepoAdapter struct {
	userDao database.UserDaoPorts
}

func NewUserRepoAdapter(dao database.UserDaoPorts) *UserRepoAdapter {
	return &UserRepoAdapter{
		userDao: dao,
	}
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

func (userRepo *UserRepoAdapter) Update(user *user.User) (*user.User, error) {
	newUser, _ := userRepo.userDao.UpdateUser(user)
	return newUser, nil
}

func (userRepo *UserRepoAdapter) Delete(id string) error {
	err := userRepo.userDao.DeleteUser(id)
	return err
}
