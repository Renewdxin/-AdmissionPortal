package vaidate

import (
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/asaskevich/govalidator"
	"log"
	"time"
)

type Validator struct {
	redisClient database.RedisPorts
}

func NewValidator(redisClient database.RedisPorts) *Validator {
	return &Validator{
		redisClient: redisClient,
	}
}

func (v Validator) EmailValidate(email string) bool {
	if !govalidator.IsExistingEmail(email) {
		log.Fatalln("INVALID EMAIL")
		return false
	} else {
		return true
	}
}

func (v Validator) PhoneValidate(phone string) bool {
	if !govalidator.IsNumeric(phone) || !govalidator.StringLength(phone, "8", "13") {
		log.Printf("INVALID PHONE NUMBER: %v", phone)
		return false
	} else {
		return true
	}
}

func (v Validator) PasswordValidate(password string) bool {
	if !govalidator.StringLength(password, "8", "20") {
		log.Println("INVALID PASSWORD")
		return false
	} else {
		return true
	}

}

func (v Validator) NameValidate(name string) bool {
	if !govalidator.StringLength(name, "2", "40") {
		log.Fatalln("INVALID LENGTH")
		return false
	}

	validNameRegex := "^[A-Za-z ]+$"
	if !govalidator.Matches(name, validNameRegex) {
		log.Fatalln("INVALID NAME")
		return false
	}

	return true
}

func (v Validator) CodeValidate(code, phone string) bool {
	return true
}

func (v Validator) BirthValidate(birth string) bool {
	// Parse the date string using the time package
	_, err := time.Parse("2006-01-02", birth)
	return err == nil
}
