package web

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(uhandler *UserHandler) *gin.Engine {
	r := gin.New()
	// home page
	r.POST("/home")

	// auth setting
	apiAccount := r.Group("/accounts")
	apiAccount.Use(middleware.JWTHandler())
	{
		apiAccount.POST("/login")
		apiAccount.POST("/logout")
		apiAccount.POST("/signup")
		apiAccount.POST("/password/forget/:id")
		apiAccount.POST("/password/change/:id")
	}

	// personal info
	apiPofile := r.Group("/profile")
	apiPofile.Use(middleware.JWTHandler())
	{
		apiPofile.GET("/Info/:id", uhandler.GetUserInfo)
		apiPofile.DELETE("/delete/:id", uhandler.DeleteUser)
		apiPofile.PUT("/update/:id", uhandler.UpdateUserInfo)
		apiPofile.GET("/status/:id", uhandler.GetUserStatus)
	}

	// recruitment
	//apiRecruitment := r.Group("/recruitment")
	//{
	//	apiRecruitment.GET("/jobs", func(c *gin.Context) {
	//	jobs, err := recruiter.ListJobs()
	//	// 处理结果并返回响应
	//})
	//
	//	apiRecruitment.GET("/jobs/:jobID", func(c *gin.Context) {
	//	jobID := c.Param("jobID")
	//	job, err := recruiter.GetJobDetails(jobID)
	//	// 处理结果并返回响应
	//})
	//
	//	apiRecruitment.POST("/apply/:jobID", middleware.JWTHandler(), func(c *gin.Context) {
	//	jobID := c.Param("jobID")
	//	userID := c.GetString("userID")
	//	err := recruiter.ApplyForJob(userID, jobID)
	//	// 处理结果并返回响应
	//})
	//}
	return r
}
