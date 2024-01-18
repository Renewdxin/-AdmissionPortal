package user

type Admin struct {
	ID       string
	password string
}

type AdminCorePorts interface {
	TableName() string
}
