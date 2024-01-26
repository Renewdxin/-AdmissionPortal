package job

type Job struct {
	ID   string `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

type JobsCorePorts interface {
	CreateJob(id, name string) Job
	TableName() string
	ShowJob(job Job) (id, name string)
}
