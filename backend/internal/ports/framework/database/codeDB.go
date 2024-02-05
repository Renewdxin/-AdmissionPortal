package database

type CodeDBPort interface {
	SaveCode(code string, email string) error
}
