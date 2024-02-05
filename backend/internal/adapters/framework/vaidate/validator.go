package vaidate

import (
	"github.com/Renewdxin/selfMade/internal/ports/framework/database"
	"github.com/asaskevich/govalidator"
	"log"
	"time"
)

type ValidatorAdapter struct {
	redisClient database.RedisDBPort
}

func NewValidatorAdapter(redisClient database.RedisDBPort) *ValidatorAdapter {
	return &ValidatorAdapter{
		redisClient: redisClient,
	}
}

func (v ValidatorAdapter) EmailValidate(email string) bool {
	if !govalidator.IsExistingEmail(email) {
		log.Fatalln("INVALID EMAIL")
		return false
	} else {
		return true
	}
}

func (v ValidatorAdapter) PhoneValidate(phone string) bool {
	if !govalidator.IsNumeric(phone) || !govalidator.StringLength(phone, "8", "13") {
		log.Printf("INVALID PHONE NUMBER: %v", phone)
		return false
	} else {
		return true
	}
}

func (v ValidatorAdapter) PasswordValidate(password string) bool {
	if !govalidator.StringLength(password, "8", "20") {
		log.Println("INVALID PASSWORD")
		return false
	} else {
		return true
	}

}

func (v ValidatorAdapter) NameValidate(name string) bool {
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

func (v ValidatorAdapter) CodeValidate(code, phone string) bool {
	return true
}

func (v ValidatorAdapter) BirthValidate(birth string) bool {
	// Parse the date string using the time package
	_, err := time.Parse("2006-01-02", birth)
	return err == nil
}
