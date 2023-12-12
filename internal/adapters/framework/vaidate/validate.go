package vaidate

import (
	"fmt"
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/asaskevich/govalidator"
	"log"
)

type Validator struct {
	redisClient database.RedisPorts
}

func NewValidator(redisClient database.RedisPorts) *Validator {
	return &Validator{
		redisClient: redisClient,
	}
}

func (v *Validator) EmailValidate(email string) bool {
	if !govalidator.IsExistingEmail(email) {
		log.Fatalln("INVALID EMAIL")
		return false
	} else {
		return true
	}
}

func (v *Validator) PhoneValidate(phone string) bool {
	if !govalidator.IsNumeric(phone) || !govalidator.StringLength(phone, fmt.Sprintf("%d, %d", 8, 13)) {
		log.Fatalln("INVALID PHONE NUMBER")
		return false
	} else {
		return true
	}
}

func (v *Validator) PasswordValidate(password string) bool {
	if !govalidator.StringLength(password, fmt.Sprintf("%d, %d", 6, 20)) {
		log.Fatalln("INVALID PASSWORD")
		return false
	} else {
		return true
	}

}

func (v *Validator) NameValidate(name string) bool {
	if !govalidator.StringLength(name, fmt.Sprintf("%d,%d", 2, 40)) {
		log.Fatalln("INVALID NAME")
		return false
	}

	validNameRegex := "^[A-Za-z ]+$"
	if !govalidator.Matches(name, validNameRegex) {
		log.Fatalln("INVALID NAME")
		return false
	}

	return true
}
