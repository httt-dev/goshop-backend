package router

import (
	"github.com/labstack/echo/v4"
	"timewise/handler"
	"timewise/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	CateHandler handler.CateHandler
}

func (api *API) SetupRouter() {
	//user API
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandlerSignUp)
	user.POST("/sign-in", api.UserHandler.HandlerSignIn)
	//user.GET("/profile/:id", api.UserHandler.HandlerProfile , middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.HandlerProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandlerListUsers , middleware.JWTMiddleware())

	//categories
	categories := api.Echo.Group("/cate", middleware.JWTMiddleware() , middleware.CheckAdminRole())
	categories.POST("/add", api.CateHandler.HandlerAddCate)
	categories.PUT("/edit", api.CateHandler.HandlerEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandlerCateDetail)
	categories.GET("/list", api.CateHandler.HandlerCateList)
}

func (api *API) SetupAdminRouter()  {
	//admin route
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenAdminToken)
	admin.POST("/sign-in", api.UserHandler.HandlerAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandlerAdminSignUp, middleware.JWTMiddleware())
}
