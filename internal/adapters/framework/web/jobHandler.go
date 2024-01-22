package web

import (
	"encoding/json"
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
	result := adapter.app.ShowAllJobs()

	// 使用循环遍历结构体数组
	for _, data := range result {
		// 将每个结构体转换为 JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// 输出 JSON 数据
		c.String(http.StatusOK, string(jsonData))
	}
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
