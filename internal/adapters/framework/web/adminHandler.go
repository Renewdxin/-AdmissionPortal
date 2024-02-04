package web

import (
	"encoding/json"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/app/job"
	"github.com/Renewdxin/selfMade/internal/ports/app/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandlerAdapter struct {
	AdminApp user.AdminApplicationPort
	JobApp   job.JobsApplicationPort
	UserApp  user.UsrApplicationPort
}

func NewAdminHandlerAdapter(AdminApp user.AdminApplicationPort, JobApp job.JobsApplicationPort, UserApp user.UsrApplicationPort) AdminHandlerAdapter {
	return AdminHandlerAdapter{
		AdminApp: AdminApp,
		JobApp:   JobApp,
		UserApp:  UserApp,
	}
}

func (adapter AdminHandlerAdapter) HomePage(c *gin.Context) {
	// 传递给模板的动态数据
	data := gin.H{
		"title": "Admin Homepage",
	}

	// 渲染HTML页面，使用已有的HTML文件 "your_existing_template.html"
	c.HTML(http.StatusOK, "your_existing_template.html", data)
}

func (adapter AdminHandlerAdapter) ShowJobApply(c *gin.Context) {
	result := adapter.AdminApp.ShowJobsApply()

	// 使用循环遍历结构体数组
	for _, index := range result {
		// 将每个结构体转换为 JSON
		jsonData, err := json.Marshal(index)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		// 输出 JSON 数据
		c.String(http.StatusOK, string(jsonData))
	}
}

func (adapter AdminHandlerAdapter) ShowAllJobs(c *gin.Context) {
	result := adapter.AdminApp.ShowAllJobs()

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

func (adapter AdminHandlerAdapter) ShowJobDetails(c *gin.Context) {
	id := c.GetHeader("id")
	details := adapter.JobApp.FindJobByID(id)
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

func (adapter AdminHandlerAdapter) ApproveJobs(c *gin.Context) {
	id := c.GetHeader("id")
	details, err := adapter.UserApp.GetUserProfile(id)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "AdminHandlerAdapter ApproveJobs GetUserInfo Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	if !adapter.AdminApp.ApproveJobs(details) {
		logger.Logger.Logf(logger.ErrorLevel, "AdminHandlerAdapter ApproveJobs ApproveJobs Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Change",
	})
}
