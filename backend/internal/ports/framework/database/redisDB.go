package database

type RedisDBPort interface {
	CloseConnection() error
	SaveVerificationCode(email, code string) error
	GetVerificationCode(email string) (string, error)
}
