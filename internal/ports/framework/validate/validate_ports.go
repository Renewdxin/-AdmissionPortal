package validate

type Validator interface {
	EmailValidate(email string) bool
	PhoneValidate(phone string) bool
	PasswordValidate(password string) bool
	NameValidate(name string) bool
	CodeValidate(code, phone string) bool
}
