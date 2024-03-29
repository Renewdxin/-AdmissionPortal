package validate

type ValidatorPort interface {
	EmailValidate(email string) bool
	PhoneValidate(phone string) bool
	PasswordValidate(password string) bool
	NameValidate(name string) bool
	CodeValidate(code, phone string) bool
	BirthValidate(birth string) bool
}
