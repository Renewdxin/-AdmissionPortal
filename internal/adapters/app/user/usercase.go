package user

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/core/verifycation"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	"log"
)

type API struct {
	userService user.RepositoryPorts
	userDao     database.UserDaoPorts
	mailSender  mail.MailPorts
	//validator    validate.Validator
	verify      verifycation.CodePorts
	redisClient database.RedisPorts
}

func NewUserAPI(usrService user.RepositoryPorts, userDao database.UserDaoPorts, mailSender mail.MailPorts,
	verify verifycation.CodePorts, redisClient database.RedisPorts) *API {
	return &API{
		userService: usrService,
		userDao:     userDao,
		mailSender:  mailSender,

		verify:      verify,
		redisClient: redisClient,
	}
}

func (api API) IfExist(email string) bool {
	return api.userDao.IfExist(email)
}

func (api API) RegisterUser(name, gender, email, phone string) (*user.User, error) {
	// if already exists
	if api.IfExist(email) {
		return nil, errors.New("user already exists")
	}

	// code generate and save
	var err error
	code := api.verify.GenerateCode()
	err = api.redisClient.SaveVerificationCode(email, code)
	if err != nil {
		log.Fatalf("Failed to generate the code, plz try again")
		return nil, err
	}
	//code send
	err = api.mailSender.CodeSend(email, "Verify your email", code)
	if err != nil {
		log.Fatalf("Failed to send code, plz try again")
		return nil, err
	}
	// verify code
	verify, _ := api.redisClient.GetVerificationCode(email)
	if verify != code {
		log.Fatalln("INVALID CODE")
		return nil, err
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

func (api API) GetUserProfile(id string) (*user.User, error) {
	// if already exists
	return api.userService.FindByID(id)
}
