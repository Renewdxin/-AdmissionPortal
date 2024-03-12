package main

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/auth"
	jobApp "github.com/Renewdxin/selfMade/internal/adapters/app/job"
	"github.com/Renewdxin/selfMade/internal/adapters/app/middleware"
	"github.com/Renewdxin/selfMade/internal/adapters/app/user"
	authApp "github.com/Renewdxin/selfMade/internal/adapters/core/auth"
	"github.com/Renewdxin/selfMade/internal/adapters/core/job"
	userApp "github.com/Renewdxin/selfMade/internal/adapters/core/user"
	"github.com/Renewdxin/selfMade/internal/adapters/core/verify"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/database"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/mail"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/vaidate"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/web"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var (
	authHandler  *web.AuthHandlerAdapter
	homeHandler  *web.HomeHandlerAdapter
	jwtHandler   *web.JWTHandlerAdapter
	adminHandler *web.AdminHandlerAdapter
	userHandler  *web.UsrHandlerAdapter
	jobHandler   *web.JobHandlerAdapter
)

func init() {
	NewLogger()
	LoadENV()
	redisClient := database.NewRedisClientAdapter()

	mailSender := mail.NewMailAdapter()
	verification := verify.NewVerificationCodeAdapter()
	validator := vaidate.NewValidatorAdapter(redisClient)
	messageSender, err := mail.NewSMSAdapter(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
	if err != nil {
		logger.Logger.Logf(logger.FatalLevel, "无法加载 messageSender, 错误信息：  %v", err)
	}

	jwtAPI := middleware.NewJWTAdapters()
	jwtHandler = web.NewJWTHandlerAdapter(jwtAPI)

	userCore := userApp.NewUsrCoreAdapter()
	userDao, err := database.NewUsrDaoAdapter(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), userCore)
	if err != nil {
		logger.Logger.Log(logger.FatalLevel, "failed to connect to the user database")
	}

	authCore := authApp.NewAuthorizeCoreAdapter()
	authDao, err := database.NewAuthDaoAdapter(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), authCore)
	if err != nil {
		logger.Logger.Log(logger.FatalLevel, "failed to connect to the account database")
	}

	userAPI := user.NewUsrApplicationAdapter(userCore, userDao, mailSender, verification, redisClient, validator)
	userHandler = web.NewUserHandlerAdapter(userAPI)

	authAPI := auth.NewAuthCaseAdapter(authCore, authDao, verification, redisClient, validator, mailSender)
	authHandler = web.NewAuthHandlerAdapter(authAPI, jwtAPI)

	homeHandler = web.NewHomeHandlerAdapter()

	jobCore := job.NewJobsCoreAdapter()
	jobDao := database.NewJobsDaoAdapter(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), jobCore)
	jobAPI := jobApp.NewJobCaseAdapter(jobCore, jobDao, userDao)
	jobHandler = web.NewJobHandlerAdapter(jobAPI)

	adminCore := userApp.NewAdminCoreAdapter()
	_ = database.NewAdminDaoAdapter(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), adminCore)
	adminAPI := user.NewAdminAppAdapter(jobAPI, jobDao, userDao)
	adminHandler = web.NewAdminHandlerAdapter(adminAPI, jobAPI, userAPI, messageSender)
}

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logger.Logf(logger.FatalLevel, "无法加载 .env 文件: %v", err)
	}
}

func NewLogger() {
	logger.Logger = logger.NewLogAdapter()
}

func InitializeRouter() {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // 允许的前端域
	r.Use(cors.New(config))

	// 注册路由和中间件
	RegisterRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		logger.Logger.Logf(logger.FatalLevel, "failed to start : %v", err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	// 注册 home 页面路由
	RegisterHomePage(r)

	// 注册 auth 相关路由
	RegisterAuthRoutes(r)

	// 注册 profile 相关路由
	RegisterProfileRoutes(r)

	// 注册 recruitment 相关路由
	RegisterRecruitmentRoutes(r)

	// 注册 admin 相关路由
	RegisterAdminRoutes(r)
}

func RegisterHomePage(r *gin.Engine) {
	homeHandler := web.NewHomeHandlerAdapter()
	r.POST("/home", homeHandler.HomePage)
}

func RegisterAuthRoutes(r *gin.Engine) {
	apiAccount := r.Group("/auth")
	apiAccount.Use(jwtHandler.JWTHandler())
	{
		apiAccount.POST("/login", authHandler.Login)
		apiAccount.POST("/signup", authHandler.Register)
		apiAccount.POST("/code/send", authHandler.CodeSend)
		apiAccount.POST("/code/verify", authHandler.CodeVerify)
		apiAccount.POST("/password/change/code", authHandler.ChangePasswordByCode)
		apiAccount.POST("/password/change/pwd", authHandler.ChangePasswordByPwd)
	}
}

func RegisterProfileRoutes(r *gin.Engine) {
	apiProfile := r.Group("/profile")
	apiProfile.Use(jwtHandler.JWTHandler())
	{
		// GetUserInfo 通过id得到用户信息，返回项为姓名、性别、出生日期、邮箱、手机号
		apiProfile.GET("/Info/:id", userHandler.GetUserInfo)
		// DeleteUser 用户在删除前需要进行手机验证码验证才能删除
		apiProfile.DELETE("/delete/:id", userHandler.DeleteUser)
		// 更新用户信息，仅限手机号、邮箱
		apiProfile.PUT("/update/:id", userHandler.UpdateUserInfo)
		// 查询是否被录取
		apiProfile.GET("/status/:id", userHandler.GetUserStatus)
	}
}

func RegisterRecruitmentRoutes(r *gin.Engine) {

	apiJob := r.Group("/recruitment")
	apiJob.Use()
	{
		//查看岗位总览
		apiJob.GET("/jobs", jobHandler.GetJobs)
		//查看岗位详细信息
		apiJob.GET("/job/:jobID", jobHandler.GetJobInfo)
		//申请投递
		apiJob.POST("/job/:jobID/apply/:userID", jobHandler.ApplyJob)
	}

}

func RegisterAdminRoutes(r *gin.Engine) {
	apiAdmin := r.Group("/admin")
	apiAdmin.Use() // 使用JWT中间件进行管理员身份验证
	{
		// 管理员仪表板或主页
		apiAdmin.GET("/dashboard", adminHandler.HomePage)

		// 查看所有职位发布（管理员）
		apiAdmin.GET("/jobs", adminHandler.ShowAllJobs)

		// 查看职位详情（管理员）
		apiAdmin.GET("/job/:jobID", adminHandler.ShowJobDetails)

		// 查看职位申请（管理员）
		apiAdmin.GET("/applications/:jobID", adminHandler.ShowJobApply)
		apiAdmin.GET("/applications/all", adminHandler.ShowAllJobApply)

		// 获取未处理的人
		apiAdmin.GET("/applications/unhandled", adminHandler.ShowAllUnhandledApply)
		apiAdmin.GET("/applications/unhandled/:jobID", adminHandler.ShowUnhandledApply)

		// 审批或拒绝职位申请（管理员）
		apiAdmin.PUT("/application/approve", adminHandler.ApproveJobs)
	}
}

func main() {
	// 初始化路由和中间件
	InitializeRouter()
}
