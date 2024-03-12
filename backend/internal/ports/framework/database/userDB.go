package database

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type UserDBPort interface {
	SaveUser(user user.User) error
	DeleteUser(id string) error
	UpdateUser(user.User) error
	FindUserByID(id string) (user.User, error)
	FindUserByEmail(email string) (user.User, error)
	IfExist(email string) bool
	ChangeUserStatus(id string, state int) bool
	ShowAllUsers() ([]user.User, error)
	ShowJobApplyByJobID(jobID string) ([]user.User, error)
	ShowAllUnhandledApplications() ([]user.User, error)
	ShowUnhandledApplicationsByJobID(jobID string) ([]user.User, error)
}
