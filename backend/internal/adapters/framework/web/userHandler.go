package web

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	userAPI "github.com/Renewdxin/selfMade/internal/ports/app/user"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UsrHandlerAdapter struct {
	userCase userAPI.UsrApplicationPort
}

func NewUserHandlerAdapter(userCase userAPI.UsrApplicationPort) *UsrHandlerAdapter {
	return &UsrHandlerAdapter{
		userCase: userCase,
	}
}

// GetUserInfo get user info
func (Handler UsrHandlerAdapter) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "server error",
		})
	}
	newUser, err := Handler.userCase.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "please try again",
		})
		log.Fatalf("No such user: %v", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":         "Your profile",
		"name":        newUser.Name,
		"gender":      newUser.Gender,
		"birth":       newUser.Birth,
		"email":       newUser.Email,
		"phoneNumber": newUser.PhoneNumber,
	})

}

// GetUserStatus be hired or not
func (Handler UsrHandlerAdapter) GetUserStatus(c *gin.Context) {
	userID := c.Param("id")
	newUser, err := Handler.userCase.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "please try again",
		})
		log.Fatalf("No such user: %v", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "Your profile",
		"state": newUser.State,
	})
}

// UpdateUserInfo update user info
func (Handler UsrHandlerAdapter) UpdateUserInfo(c *gin.Context) {
	var newUser user.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "User :%v Bind Error", newUser.ID)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "server error",
		})
	}
	if err := Handler.userCase.UpdateUser(newUser); err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "User :%v Update Error", newUser.ID)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Request Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// DeleteUser delete user
func (Handler UsrHandlerAdapter) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	if err := Handler.userCase.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "please try again",
		})
		logger.Logger.Logf(logger.ErrorLevel, "User :%v Delete Error", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "SUCCESS",
	})
}
