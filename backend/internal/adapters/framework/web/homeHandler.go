package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeHandlerAdapter struct {
}

func NewHomeHandlerAdapter() *HomeHandlerAdapter {
	return &HomeHandlerAdapter{}
}

func (adapter HomeHandlerAdapter) HomePage(c *gin.Context) {
	// 传递给模板的动态数据
	data := gin.H{
		"title": "Home",
	}

	// 渲染HTML页面，使用已有的HTML文件 "your_existing_template.html"
	c.HTML(http.StatusOK, "your_existing_template.html", data)
}
