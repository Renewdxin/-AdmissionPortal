package main

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/user"
	user2 "github.com/Renewdxin/selfMade/internal/adapters/core/user"
	"github.com/Renewdxin/selfMade/internal/adapters/core/verify"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/database"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/mail"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/vaidate"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/web"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	redisClient := database.NewRedisClient()
	userDao := database.NewUserDao(db)
	mailSender := mail.NewMail()
	verification := verify.NewVerificationCodeService()
	validator := vaidate.NewValidator(redisClient)
	userCore := user2.NewUserService(validator)
	userAPI := user.NewUserAPI(userCore, userDao, mailSender, verification, redisClient)
	uHandler := web.NewUserHandler(userAPI)
	route := web.NewRouter(uHandler)

	route.Run("8080")
}
