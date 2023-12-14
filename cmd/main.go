package main

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/auth"
	"github.com/Renewdxin/selfMade/internal/adapters/app/user"
	auth2 "github.com/Renewdxin/selfMade/internal/adapters/core/auth"
	user2 "github.com/Renewdxin/selfMade/internal/adapters/core/user"
	"github.com/Renewdxin/selfMade/internal/adapters/core/verify"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/database"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/utils/mail"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/utils/vaidate"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("internal/adapters/framework/global/.env")
	if err != nil {
		log.Fatalf("无法加载 .env 文件: %v", err)
	}

	redisClient := database.NewRedisClient()

	mailSender := mail.NewMail()
	verification := verify.NewVerificationCodeService()
	validator := vaidate.NewValidator(redisClient)
	userCore := user2.NewUserService()
	userDao, err := database.NewUserDao(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), userCore)
	if err != nil {
		log.Fatalf("falied to connect to the user database")
	}

	authCore := auth2.NewAdapter()
	authDao, err := database.NewauthDao(os.Getenv("DRIVER_NAME"), os.Getenv("DRIVER_SOURCE_NAME"), authCore)
	if err != nil {
		log.Fatalf("falied to connect to the user database")
	}
	userAPI := user.NewUserAPI(userCore, userDao, mailSender, verification, redisClient, validator)
	userHandler := web.NewUserHandler(userAPI)

	authAPI := auth.NewAuthCaseAdapter(authCore, authDao, verification, redisClient, validator, mailSender)
	authHandler := web.NewAuthHandler(authAPI)

	r := gin.New()
	// home page
	//r.POST("/home")

	// auth setting
	apiAccount := r.Group("/auth")
	apiAccount.Use()
	{
		apiAccount.POST("/login", authHandler.Login)
		//apiAccount.POST("/logout")
		apiAccount.POST("/signup", authHandler.Register)
		apiAccount.POST("/password/forget/:id", authHandler.ForgetPassword)
		apiAccount.POST("/password/change/:id", authHandler.ChangePassword)
	}

	// personal info
	apiProfile := r.Group("/profile")
	apiProfile.Use()
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
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("falied to start : %v", err)
	}
}
