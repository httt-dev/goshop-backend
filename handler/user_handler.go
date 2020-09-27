package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"timewise/model"
	"timewise/repository"
	"timewise/security"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u UserHandler) HandlerSignUp(c echo.Context) error {

	req := model.SignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//
	/*if err:= c.Validate(req);err !=nil{
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode : http.StatusBadRequest,
			Message : err.Error(),
			Data : nil,
		})
	}*/

	//
	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//
	user := model.User{
		UserID:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		Password: hash,
		Role:     role,
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//get token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})

}

func (u UserHandler) HandlerSignIn(c echo.Context) error {
	req := model.SignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user, err := u.UserRepo.SelectUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	//check password
	isTheSame := security.ComparePasswords(user.Password, []byte (req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}
	//get token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công",
		Data:       user,
	})

}

func (u UserHandler) HandlerProfile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	//userId:=c.Param("id")
	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công",
		Data:       user,
	})
}

/**
Lấy danh sách users
 */
func (u UserHandler) HandlerListUsers( c echo.Context) error {
	users , err := u.UserRepo.SelectUsers(c.Request().Context())
	if err !=nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data : nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Xử lý thành công",
		Data:users,
	})
}