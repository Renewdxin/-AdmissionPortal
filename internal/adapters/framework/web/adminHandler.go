package web

import (
	"github.com/Renewdxin/selfMade/internal/ports/app/user"
	"github.com/gin-gonic/gin"
)

type AdminHandlerAdapter struct {
	app user.AdminApplicationPorts
}

func NewAdminHandlerAdapter() AdminHandlerAdapter {
	return AdminHandlerAdapter{}
}

func (adapter AdminHandlerAdapter) HomePage() {

}

func (adapter AdminHandlerAdapter) ShowJobApply(c *gin.Context) {

}

func (adapter AdminHandlerAdapter) ShowAllJobs(c *gin.Context) {

}

func (adapter AdminHandlerAdapter) ShowJobDetails(c *gin.Context) {

}

func (adapter AdminHandlerAdapter) ApproveJobs(c *gin.Context) {

}
