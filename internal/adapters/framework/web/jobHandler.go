package web

import (
	"github.com/Renewdxin/selfMade/internal/ports/app/job"
	job2 "github.com/Renewdxin/selfMade/internal/ports/core/job"
	user2 "github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobHandlerAdapter struct {
	app job.JobsCasePorts
}

func NewAdapter(app job.JobsCasePorts) JobHandlerAdapter {
	return JobHandlerAdapter{
		app: app,
	}
}

func (adapter JobHandlerAdapter) GetJobs(c *gin.Context) {

}

func (adapter JobHandlerAdapter) GetJobInfo(c *gin.Context) {
	id := c.GetHeader("jobID")
	jobDetail := adapter.app.FindJobByID(id)
	c.JSON(http.StatusOK, gin.H{
		"id":   jobDetail.ID,
		"name": jobDetail.Name,
	})
}

func (adapter JobHandlerAdapter) DeleteJob(c *gin.Context) {
	id := c.GetHeader("jobID")
	jobDetail := adapter.app.FindJobByID(id)
	if adapter.app.DeleteJob(jobDetail) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Successful",
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "Error! Please try again later",
	})
}

func (adapter JobHandlerAdapter) UpdateJob(c *gin.Context) {
	var newJob job2.Job
	if err := c.ShouldBind(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error! Please try again later",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Successful",
	})
}

func (adapter JobHandlerAdapter) AddJob(c *gin.Context) {
	// 表单信息输入验证
	var user user2.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Request error",
		})
	}
	// 上传到数据库中

	// 通知admin

}
