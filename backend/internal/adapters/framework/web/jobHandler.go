package web

import (
	"encoding/json"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/app/job"
	jobCore "github.com/Renewdxin/selfMade/internal/ports/core/job"
	userCore "github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobHandlerAdapter struct {
	app job.JobsApplicationPort
}

func NewJobHandlerAdapter(app job.JobsApplicationPort) *JobHandlerAdapter {
	return &JobHandlerAdapter{
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
			logger.Logger.Log(logger.ErrorLevel, "GetJobs Marshal Error")
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
		return
	}
	logger.Logger.Log(logger.ErrorLevel, "Delete Job Error")
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "Error! Please try again later",
	})
}

func (adapter JobHandlerAdapter) UpdateJob(c *gin.Context) {
	var newJob jobCore.Job
	if err := c.ShouldBind(&newJob); err != nil {
		logger.Logger.Log(logger.ErrorLevel, "Update Job Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error! Please try again later",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Successful",
	})
}

func (adapter JobHandlerAdapter) ApplyJob(c *gin.Context) {
	// 表单信息输入验证
	var user userCore.User

	if err := c.ShouldBind(&user); err != nil {
		logger.Logger.Log(logger.ErrorLevel, "Apply Job Binding error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Request error",
		})
		return
	}

	// 上传到数据库中
	if !adapter.app.ApplyJob(user, user.ApplyID, user.ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
