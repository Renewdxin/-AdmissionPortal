package user

type Admin struct {
	ID       string
	password string
}

type AdminCorePort interface {
	TableName() string
}
