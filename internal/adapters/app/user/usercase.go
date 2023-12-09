package user

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/core/verifycation"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	"github.com/Renewdxin/selfMade/internal/ports/framework/validate"
	"log"
)

type API struct {
	userService user.RepositoryPorts
	userDao     database.UserDaoPorts
	mailSender  mail.MailPorts
	validator   validate.Validator
	code        verifycation.CodePorts
}

func NewUserAPI(usrService user.RepositoryPorts, userDao database.UserDaoPorts, mailSender mail.MailPorts,
	validator validate.Validator, code verifycation.CodePorts) *API {
	return &API{
		userService: usrService,
		userDao:     userDao,
		mailSender:  mailSender,
		validator:   validator,
		code:        code,
	}
}

func (api API) IfExist(email string) bool {
	return api.userDao.IfExist(email)
}

func (api API) RegisterUser(name, gender, email, phone string) (*user.User, error) {
	// if already exists
	if api.IfExist(email) {
		code := api.code.GenerateCode()
		// code send
		err := api.mailSender.CodeSend(email, "Verify your email", code)
		if err != nil {
			log.Fatalf("Failed to send emails, plz try again")
			return nil, err
		}
		// code validate
		if !api.validator.CodeValidate(email, code) {
			return nil, errors.New("wrong code")
		}
		// persistence
		newUser, err := api.userDao.SaveUser(name, gender, email, phone)
		if err != nil {
			log.Fatalf("Failed to save user, plz try again")
			return nil, err
		}
		// send notification
		if err := api.mailSender.WelcomeMail(name); err != nil {
			log.Println("Failed to send welcome email")
		}
		return newUser, nil
	}
	return nil, errors.New("user already exists")

}

func (api API) GetUserProfile(id string) (*user.User, error) {
	// if already exists
	return api.userService.FindByID(id)
}
