package web

import "github.com/gin-gonic/gin"

type AdminHandlerPort interface {
	HomePage()
	ShowJobApply(c *gin.Context)
	ShowAllJobs(c *gin.Context)
	ShowJobDetails(c *gin.Context)
	ApproveJobs(c *gin.Context)
}
