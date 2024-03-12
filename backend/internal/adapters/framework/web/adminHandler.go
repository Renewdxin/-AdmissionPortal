package web

import (
	"encoding/json"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/app/job"
	"github.com/Renewdxin/selfMade/internal/ports/app/user"
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type AdminHandlerAdapter struct {
	AdminApp user.AdminApplicationPort
	JobApp   job.JobsApplicationPort
	UserApp  user.UsrApplicationPort
	Notify   mail.SMSPorts
}

func NewAdminHandlerAdapter(AdminApp user.AdminApplicationPort, JobApp job.JobsApplicationPort, UserApp user.UsrApplicationPort, Notify mail.SMSPorts) *AdminHandlerAdapter {
	return &AdminHandlerAdapter{
		AdminApp: AdminApp,
		JobApp:   JobApp,
		UserApp:  UserApp,
		Notify:   Notify,
	}
}

func (handler AdminHandlerAdapter) HomePage(c *gin.Context) {
	// 传递给模板的动态数据
	data := gin.H{
		"title": "Admin Homepage",
	}

	// 渲染HTML页面，使用已有的HTML文件 "your_existing_template.html"
	c.HTML(http.StatusOK, "your_existing_template.html", data)
}

func (handler AdminHandlerAdapter) ShowJobApply(c *gin.Context) {
	result, err := handler.AdminApp.ShowJobsApply()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func (handler AdminHandlerAdapter) ShowAllJobApply(c *gin.Context) {
	jobID := c.Param("jobID") // 从URL参数中获取岗位ID
	result, err := handler.AdminApp.ShowJobApplyByJobID(jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, jsonData)
}

func (handler AdminHandlerAdapter) ShowAllJobs(c *gin.Context) {
	result := handler.AdminApp.ShowAllJobs()

	// 使用循环遍历结构体数组
	for _, data := range result {
		// 将每个结构体转换为 JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Logf(logger.ErrorLevel, "Admin_ShowAllJobs : Marshall Error: %v", data)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error! Please Try Again"})
			return
		}
		// 输出 JSON 数据
		c.String(http.StatusOK, string(jsonData))
	}
}

func (handler AdminHandlerAdapter) ShowJobDetails(c *gin.Context) {
	id := c.GetHeader("id")
	details := handler.JobApp.FindJobByID(id)
	if details.Name == "" {
		logger.Logger.Logf(logger.ErrorLevel, "AdminHandlerAdapter ShowJobDetails Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": details.Name,
	})
}

func (handler AdminHandlerAdapter) ApproveJobs(c *gin.Context) {
	id := c.GetHeader("id")
	// 确认用户是否存在
	details, err := handler.UserApp.GetUserProfile(id)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "AdminHandlerAdapter ApproveJobs GetUserInfo Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	// 更改用户状态
	if !handler.AdminApp.ApproveJobs(details) {
		logger.Logger.Logf(logger.ErrorLevel, "AdminHandlerAdapter ApproveJobs ApproveJobs Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	// 发送通知短信
	if err := handler.Notify.SendNotifySms(details.PhoneNumber, os.Getenv("SIGN")); err != nil {
		logger.Logger.Logf(logger.InfoLevel, "failed to send message to user : %v", id)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Change",
	})
}

func (handler AdminHandlerAdapter) ShowAllUnhandledApply(c *gin.Context) {
	result, err := handler.AdminApp.ShowAllUnhandledApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler AdminHandlerAdapter) ShowUnhandledApply(c *gin.Context) {
	jobID := c.Param("jobID")
	result, err := handler.AdminApp.ShowUnhandledApplicationsByJobID(jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, result)
}
