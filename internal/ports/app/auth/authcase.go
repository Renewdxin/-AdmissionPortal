package auth

type AuthcasePorts interface {
	BeforeRegister(id, password string) error
	RegisterByEmail(email, password string) error
	ForgetPasswordByEmail(email, password string) error
	ChangePassword(id, oldPassword, newPassword string) error
	LogIn(id, password string) error
}
