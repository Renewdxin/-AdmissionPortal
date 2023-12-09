package database

type CodePorts interface {
	SaveCode(code string, email string) error
}
