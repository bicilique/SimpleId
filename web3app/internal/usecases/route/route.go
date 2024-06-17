package route

import (
	"SimpleId/internal/handlers"
	"SimpleId/internal/services"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Api          *gin.Engine
	UserHandler  *handlers.UserHandler
	AdminHandler *handlers.AdminHandler
}

func (c *RouteConfig) Setup() {
	c.setupUser()
	c.setupadmin()
	c.setupPublic()
}

func (c *RouteConfig) setupPublic() {
	public := c.Api.Group("/api")
	public.POST("/register", c.UserHandler.Register)
	public.POST("/login", c.UserHandler.Login)
}
func (c *RouteConfig) setupUser() {
	user := c.Api.Group("/api/user")
	user.Use(services.JWTAuthCustomer())

	user.POST("/information", c.UserHandler.GetUserInformation)
	user.POST("/add/info", c.UserHandler.AddUserInformation)
	user.POST("/update/info", c.UserHandler.UpdateUserInformation)
}

func (c *RouteConfig) setupadmin() {
	admin := c.Api.Group("/api/admin")

	admin.Use(services.JWTAuth())
	admin.POST("/request", c.AdminHandler.RequestSharing)
	admin.POST("/approve", c.AdminHandler.ApproveUser)

	admin.POST("/sharing", c.AdminHandler.GetSharingRequest)
	admin.GET("/sharing/all", c.AdminHandler.ShowAllSharing)
	admin.POST("/sharing/status", c.AdminHandler.ShowAllSharingByStatus)
	admin.POST("/sharing/approve", c.AdminHandler.ApproveSharing)
	admin.POST("/sharing/receive", c.AdminHandler.ReceiveSharing)

}
