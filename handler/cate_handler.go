package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"timewise/app_constant"
	"timewise/log"
	"timewise/model"
	"timewise/repository"
)

type CateHandler struct {
	CateRepo repository.CateRepo
}

func (c CateHandler) HandlerAddCate(ctx echo.Context) error  {
	req := model.Cate{

	}
	if err := ctx.Bind(&req);err !=nil{
		log.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	cateId,_ := uuid.NewUUID()
	req.CateId =  cateId.String()
	cate, err := c.CateRepo.SaveCate(ctx.Request().Context() , req)
	if err!=nil {
		return ctx.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       cate,
	})
}

func (c CateHandler) HandlerEditCate(ctx echo.Context) error  {
	req := model.Cate{}
	if err:= ctx.Bind(&req) ; err!=nil{
		log.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err := c.CateRepo.UpdateCate(ctx.Request().Context() , req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (c CateHandler) HandlerCateDetail(ctx echo.Context) error {
	cateId := ctx.Param("id")

	cate, err := c.CateRepo.SelectCateById(ctx.Request().Context(), cateId)
	if err != nil {
		if err == app_constant.DataNotFound {
			return ctx.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return ctx.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       cate,
	})
}

func (c CateHandler) HandlerCateList(ctx echo.Context) error {
	cates, err := c.CateRepo.SelectCates(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       cates,
	})
}