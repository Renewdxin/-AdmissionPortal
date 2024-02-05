package user

type Admin struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
}

type AdminCorePort interface {
	TableName() string
}
