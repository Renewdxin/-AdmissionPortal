package database

type RedisPorts interface {
	CloseConnection() error
	SaveVerificationCode(email, code string) error
	GetVerificationCode(email string) (string, error)
}
