package web

import "github.com/gin-gonic/gin"

type UserHandlerPort interface {
	// GetUserInfo get user info
	GetUserInfo(c *gin.Context)
	// GetUserStatus be hired or not
	GetUserStatus(c *gin.Context)
	// UpdateUserInfo update user info
	UpdateUserInfo(c *gin.Context)
	// DeleteUser delete user
	DeleteUser(c *gin.Context)
}
