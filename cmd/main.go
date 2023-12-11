package main

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/user"
	user2 "github.com/Renewdxin/selfMade/internal/adapters/core/user"
	"github.com/Renewdxin/selfMade/internal/adapters/core/verify"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/database"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/mail"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/vaidate"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/web"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	redisClient := database.NewRedisClient()

	mailSender := mail.NewMail()
	verification := verify.NewVerificationCodeService()
	validator := vaidate.NewValidator(redisClient)
	userCore := user2.NewUserService()
	userDao, err := database.NewUserDao("pgx",
		"host=localhost user=postgres password=26221030 dbname=self sslmode=disable", userCore)
	if err != nil {
		log.Fatalf("falied to connect to the user database")
	}
	userAPI := user.NewUserAPI(userCore, userDao, mailSender, verification, redisClient, validator)
	handler := web.NewUserHandler(userAPI)
	//route := web.NewRouter(uHandler)

	r := gin.New()
	// home page
	//r.POST("/home")

	// account setting
	//apiAccount := r.Group("/accounts")
	//apiAccount.Use()
	//{
	//	apiAccount.POST("/login")
	//	apiAccount.POST("/logout")
	//	apiAccount.POST("/signup")
	//	apiAccount.POST("/password/forget/:id")
	//	apiAccount.POST("/password/change/:id")
	//}

	// personal info
	apiPofile := r.Group("/profile")
	apiPofile.Use()
	{
		apiPofile.GET("/Info/:id", handler.GetUserInfo)
		apiPofile.DELETE("/delete/:id", handler.DeleteUser)
		apiPofile.PUT("/update/:id", handler.UpdateUserInfo)
		apiPofile.GET("/status/:id", handler.GetUserStatus)
	}
	r.Run(":8080")
}
