package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"timewise/app_constant"
	"timewise/model"
	"timewise/security"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.JwtKey),
	}
	return middleware.JWTWithConfig(config)
}

func CheckAdminRole() echo.MiddlewareFunc  {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			tokeData := context.Get("user").(*jwt.Token)
			claims := tokeData.Claims.(*model.JwtCustomClaims)
			if claims.Role != model.ADMIN.String() {
				return context.JSON(http.StatusForbidden, model.Response{
					StatusCode: http.StatusForbidden,
					Message:    app_constant.Http_AccessNotAllow,
					Data:       nil,
				})
			}
			return next(context)
		}

	}
}