package auth

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/Renewdxin/selfMade/internal/ports/core/verifycation"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	"github.com/Renewdxin/selfMade/internal/ports/framework/validate"
)

type AuthcaseAdapter struct {
	core        auth.AccountCorePorts
	dao         database.AuthDaoPorts
	codeGen     verifycation.CodePorts
	redisClient database.RedisPorts
	validator   validate.Validator
	mailSender  mail.MailPorts
}

func NewAuthCaseAdapter(core auth.AccountCorePorts, dao database.AuthDaoPorts,
	codeGen verifycation.CodePorts, redisClient database.RedisPorts,
	validator validate.Validator, mailSender mail.MailPorts) *AuthcaseAdapter {
	return &AuthcaseAdapter{
		core:        core,
		dao:         dao,
		codeGen:     codeGen,
		redisClient: redisClient,
		validator:   validator,
		mailSender:  mailSender,
	}
}

func (api AuthcaseAdapter) BeforeRegister(id, password string) error {
	if !api.validator.PhoneValidate(id) {
		return errors.New("INVALID NUMBER")
	}
	if !api.validator.PasswordValidate(password) {
		return errors.New("INVALID PASSWORD")
	}
	return nil
}

func (api AuthcaseAdapter) RegisterByEmail(email, password string) error {
	// validate first
	if err := api.BeforeRegister(email, password); err != nil {
		return err
	}

	account, err := api.core.CreateAccount(email, password)
	if err != nil {
		return err
	}
	// if register
	if api.dao.EmailIfExist(email) {
		return errors.New("this account already exists")
	}
	// code send

	code := api.codeGen.GenerateCode()
	err = api.redisClient.SaveVerificationCode(email, code)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "Failed to generate the code, plz try again: %v", err)
		return err
	}
	//code send
	err = api.mailSender.CodeSend(email, "Verify your email", code)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "Failed to send code, plz try again: %v", err)
		return err
	}
	// verify code
	verify, _ := api.redisClient.GetVerificationCode(email)
	if verify != code {
		logger.Logger.Logf(logger.InfoLevel, "INVALID CODE: %v", err)
		return err
	}

	if api.dao.SaveUserAccount(account) {
		return errors.New("failed to register, please try again")
	}
	return nil
}
func (api AuthcaseAdapter) ForgetPasswordByEmail(email, password string) error {
	// code generate and save to redis
	code := api.codeGen.GenerateCode()
	var err error
	err = api.redisClient.SaveVerificationCode(email, code)
	if err != nil {
		logger.Logger.Log(logger.ErrorLevel, "Failed to generate the code, plz try again")
		return err
	}
	//code send
	err = api.mailSender.CodeSend(email, "Verify your email", code)
	if err != nil {
		logger.Logger.Log(logger.ErrorLevel, "Failed to send code, plz try again")
		return err
	}
	// verify code
	verify, _ := api.redisClient.GetVerificationCode(email)
	if verify != code {
		logger.Logger.Logf(logger.InfoLevel, "INVALID CODE: %v", err)
		return err
	}
	account, errs := api.core.CreateAccount(email, password)
	if errs != nil {
		return errs
	}
	if api.dao.UpdateAccount(account) {
		return errors.New("failed to change, plz try again")
	}
	return nil

}

func (api AuthcaseAdapter) ChangePassword(id, oldPassword, newPassword string) error {
	// validate old password
	if oldPassword != api.dao.FindPasswordByID(id) {
		return errors.New("WRONG PASSWORD")
	}
	// validate new password's form
	if !api.validator.PasswordValidate(newPassword) {
		return errors.New("INVALID PASSWORD")
	}
	// Update
	account, err := api.core.CreateAccount(id, newPassword)
	if err != nil {
		return err
	}
	if api.dao.UpdateAccount(account) {
		return errors.New("failed to change, plz try again")
	}
	return nil

}

func (api AuthcaseAdapter) LogIn(id, password string) error {
	if !api.dao.FindAccountByID(id) {
		return errors.New("this id doesn't exist")
	}

	if password != api.dao.FindPasswordByID(id) {
		return errors.New("WRONG PASSWORD")
	}
	return nil
}
