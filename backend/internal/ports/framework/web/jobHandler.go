package web

import "github.com/gin-gonic/gin"

type JobsHandlerPort interface {
	GetJobs(c *gin.Context)
	GetJobInfo(c *gin.Context)
	DeleteJob(c *gin.Context)
	UpdateJob(c *gin.Context)
	ApplyJob(c *gin.Context)
}
