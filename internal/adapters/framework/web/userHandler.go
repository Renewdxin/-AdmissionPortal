package web

import (
	userAPI "github.com/Renewdxin/selfMade/internal/ports/app/user"
	"github.com/Renewdxin/selfMade/internal/ports/core/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	userCase userAPI.UserCasePorts
}

func NewUserHandler(userCase userAPI.UserCasePorts) *UserHandler {
	return &UserHandler{
		userCase: userCase,
	}
}

// GetUserInfo get user info
func (uHandler *UserHandler) GetUserInfo(c *gin.Context) {
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
func (uHandler *UserHandler) GetUserStatus(c *gin.Context) {
	userID := c.Param("id")
	newUser, err := uHandler.userCase.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "please try again",
		})
		log.Fatalf("No such user: %d", userID)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":        "Your profile",
		"created_at": newUser.CreatedAt,
		"state":      newUser.State,
	})
}

// UpdateUserInfo update user info
func (uHandler *UserHandler) UpdateUserInfo(c *gin.Context) {
	var newUser user.User
	if err := c.ShouldBind(&newUser); err != nil {
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
func (uHandler *UserHandler) DeleteUser(c *gin.Context) {
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
