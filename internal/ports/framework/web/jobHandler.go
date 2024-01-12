package web

import "github.com/gin-gonic/gin"

type JobsHandlerPorts interface {
	GetJobs(c *gin.Context)
	GetJobInfo(c *gin.Context)
	DeleteJob(c *gin.Context)
	UpdateJob(c *gin.Context)
	AddJob(c *gin.Context)
}
