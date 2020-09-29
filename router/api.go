package router

import (
	"github.com/labstack/echo/v4"
	"timewise/handler"
	"timewise/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	//user API
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandlerSignUp)
	user.POST("/sign-in", api.UserHandler.HandlerSignIn)
	//user.GET("/profile/:id", api.UserHandler.HandlerProfile , middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.HandlerProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandlerListUsers , middleware.JWTMiddleware())
}

func (api *API) SetupAdminRouter()  {
	//admin route
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenAdminToken)
	admin.POST("/sign-in", api.UserHandler.HandlerAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandlerAdminSignUp, middleware.JWTMiddleware())
}
