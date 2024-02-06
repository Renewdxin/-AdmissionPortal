package job

import "github.com/Renewdxin/selfMade/internal/ports/core/user"

type Job struct {
	ID        string      `json:"id" gorm:"id"`
	Name      string      `json:"name" gorm:"name"`
	ApplyList []user.User `json:"apply_list" gorm:"foreignKey:ApplyID"`
}

type JobsCorePort interface {
	CreateJob(id, name string) Job
	TableName() string
	ShowJob(job Job) (id, name string)
}
