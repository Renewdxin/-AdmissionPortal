package web

import (
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
func (uHandler UsrHandlerAdapter) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "please try again",
		})
	}
	newUser, err := uHandler.userCase.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "please try again",
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
func (uHandler UsrHandlerAdapter) GetUserStatus(c *gin.Context) {
	userID := c.Param("id")
	newUser, err := uHandler.userCase.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "please try again",
		})
		log.Fatalf("No such user: %v", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":        "Your profile",
		"created_at": newUser.CreatedAt,
		"state":      newUser.State,
	})
}

// UpdateUserInfo update user info
func (uHandler UsrHandlerAdapter) UpdateUserInfo(c *gin.Context) {
	var newUser user.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "INVALID PARAMS",
		})
	}
	if err := uHandler.userCase.UpdateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Request Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// DeleteUser delete user
func (uHandler UsrHandlerAdapter) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	if err := uHandler.userCase.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "please try again",
		})
		log.Fatalf("User :%v Delete Error", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "SUCCESS",
	})
}
