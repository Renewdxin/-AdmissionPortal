package user

type AdminCoreAdapter struct{}

func NewAdminCoreAdapter() AdminCoreAdapter {
	return AdminCoreAdapter{}
}

func (core AdminCoreAdapter) TableName() string {
	return "adminAccount"
}
