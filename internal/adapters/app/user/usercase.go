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
	user        user.UserPorts
	userDao     database.UserDaoPorts
	mailSender  mail.MailPorts
	verify      verifycation.CodePorts
	redisClient database.RedisPorts
	validator   validate.Validator
}

func NewUserAPI(user user.UserPorts, userDao database.UserDaoPorts, mailSender mail.MailPorts,
	verify verifycation.CodePorts, redisClient database.RedisPorts, validator validate.Validator) *API {
	return &API{
		user:        user,
		userDao:     userDao,
		mailSender:  mailSender,
		verify:      verify,
		redisClient: redisClient,
		validator:   validator,
	}
}

func (api API) IfExist(email string) bool {
	return api.userDao.IfExist(email)
}

func (api API) UserValidateBeforeRegister(name string, gender string, email string, phone string, birth string) bool {
	if !api.validator.NameValidate(name) || !api.validator.EmailValidate(email) || !api.validator.PhoneValidate(phone) {
		log.Fatalf("INVALID INFORMATION")
		return false
	}

	if gender != "male" && gender != "female" {
		log.Fatalf("INVALID INFORMATION")
		return false
	}
	if !api.validator.BirthValidate(birth) {
		log.Fatalf("INVALID INFORMATION")
		return false
	}
	return true
}

func (api API) RegisterUser(name, gender, email, phone, birth string) error {
	// validate
	if !api.UserValidateBeforeRegister(name, gender, email, phone, birth) {
		return errors.New("INVALID INFORMATION")
	}
	// if already exists
	if api.IfExist(email) {
		return errors.New("user already exists")
	}
	// code generate and save to redis
	var err error
	code := api.verify.GenerateCode()
	err = api.redisClient.SaveVerificationCode(email, code)
	if err != nil {
		log.Fatalf("Failed to generate the code, plz try again")
		return err
	}
	//code send
	err = api.mailSender.CodeSend(email, "Verify your email", code)
	if err != nil {
		log.Fatalf("Failed to send code, plz try again")
		return err
	}
	// verify code
	verify, _ := api.redisClient.GetVerificationCode(email)
	if verify != code {
		log.Fatalln("INVALID CODE")
		return err
	}
	// persistence
	newUser, err := api.user.CreateUser(name, gender, email, phone, birth)

	if err := api.userDao.SaveUser(newUser); err != nil {
		log.Fatalf("Failed to save user, plz try again")
		return err
	}
	// send notification
	if err := api.mailSender.WelcomeMail(name); err != nil {
		log.Println("Failed to send welcome email")
	}
	return nil
}

func (api API) GetUserProfile(id string) (user.User, error) {
	// if already exists
	return api.userDao.FindUserByID(id)
}

func (api API) DeleteUser(id string) error {
	//if exist
	u, err := api.userDao.FindUserByID(id)
	if err != nil {
		return err
	}
	// code generate and save to redis
	code := api.verify.GenerateCode()
	err = api.redisClient.SaveVerificationCode(u.Email, code)
	if err != nil {
		log.Fatalf("Failed to generate the code, plz try again")
		return err
	}
	//code send
	err = api.mailSender.CodeSend(u.Email, "Verify your email", code)
	if err != nil {
		log.Fatalf("Failed to send code, plz try again")
		return err
	}
	// verify code
	verify, _ := api.redisClient.GetVerificationCode(u.Email)
	if verify != code {
		log.Fatalln("INVALID CODE")
		return err
	}

	return api.userDao.DeleteUser(id)
}

func (api API) UpdateUser(user user.User) error {
	api.UserValidateBeforeRegister(user.Name, user.Gender, user.Email, user.PhoneNumber, user.Birth)
	return api.userDao.UpdateUser(user)
}
