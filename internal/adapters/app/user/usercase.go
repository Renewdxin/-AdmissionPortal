package user

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/core/verifycation"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	"log"
	"strconv"
)

type API struct {
	user        user.UserPorts
	userDao     database.UserDaoPorts
	mailSender  mail.MailPorts
	verify      verifycation.CodePorts
	redisClient database.RedisPorts
}

func NewUserAPI(user user.UserPorts, userDao database.UserDaoPorts, mailSender mail.MailPorts,
	verify verifycation.CodePorts, redisClient database.RedisPorts) *API {
	return &API{
		user:        user,
		userDao:     userDao,
		mailSender:  mailSender,
		verify:      verify,
		redisClient: redisClient,
	}
}

func (api API) IfExist(email string) bool {
	return api.userDao.IfExist(email)
}

func (api API) RegisterUser(name, gender, email, phone string) error {
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
	newUser, err := api.user.CreateUser(name, gender, email, phone)

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

func (api API) GetUserProfile(id int) (*user.User, error) {
	// if already exists
	return api.userDao.FindUserByID(strconv.Itoa(id))
}

func (api API) DeleteUser(id string) error {
	return api.userDao.DeleteUser(id)
}

func (api API) UpdateUser(user user.User) error {
	return api.userDao.UpdateUser(user)
}
