package auth

type AuthcasePorts interface {
	Register(id, password string) error
	ForgetPassword(id, phone, password string) error
	ChangePassword(id, oldPassword, newPassword string) error
	LogIn(id, password string) error
}
