package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"timewise/handler"
)

type API struct {
	Echo  *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter()  {
	api.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	user:=api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandlerSignUp)
}