package user

type Admin struct {
	ID       string
	password string
}

type adminCorePorts interface {
	TableName() string
}
