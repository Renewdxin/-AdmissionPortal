package user

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
)

type API struct {
	userService user.RepositoryPorts
	userDao     database.UserDaoPorts
}

func NewUserAPI(usrService user.RepositoryPorts, userDao database.UserDaoPorts) *API {
	return &API{
		userService: usrService,
		userDao:     userDao,
	}
}

func (api API) IfExist(email string) bool {
	return api.userDao.IfExist(email)
}

func (api API) RegisterUser(name, gender, email, phone string) (*user.User, error) {
	// if already exists
	if api.IfExist(email) {
		// send code
		// validate the email
		// send notification
		// persistence
		return api.userService.CreateUser(name, gender, email, phone)
	}
	return nil, errors.New("user already exists")

}

func (api API) GetUserProfile(id string) (*user.User, error) {
	// if already exists
	return api.userService.FindByID(id)
}
