package user

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/validate"
	"log"
)

type UserService struct {
	validator validate.Validator
}

func NewUserService(validator validate.Validator) *UserService {
	return &UserService{
		validator: validator,
	}
}

func (u *UserService) UserValidate(name string, gender string, email string, phone string) bool {
	if !u.validator.NameValidate(name) || !u.validator.EmailValidate(email) || !u.validator.PhoneValidate(phone) {
		log.Fatalf("INVALID INFORMATION")
		return false
	}

	if gender != "male" && gender != "female" {
		log.Fatalf("INVALID INFORMATION")
		return false
	}
	return true
}

func (u *UserService) CreateUser(name string, gender string, email string, phone string) (user.User, error) {
	// validate the u
	if !u.UserValidate(name, gender, email, phone) {
		return user.User{}, errors.New("INVALID INFORMATION")
	}
	newUser := user.User{
		Name:        name,
		Email:       email,
		PhoneNumber: phone,
		Gender:      gender,
	}
	return newUser, nil
}
