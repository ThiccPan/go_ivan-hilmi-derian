package handler

import (
	"net/http"

	"github.com/alta_code_competence/helper"
	"github.com/alta_code_competence/internal/domain"
	"github.com/alta_code_competence/internal/user/usecase"
	"github.com/labstack/echo/v4"
)

type HttpHandler interface{}

type httpHandler struct {
	usecase usecase.UserUsecase
	tokenManager helper.AuthJWT
}

func NewHttpUserHandler(u usecase.UserUsecase, tm helper.AuthJWT) *httpHandler {
	return &httpHandler{
		usecase: u,
		tokenManager: tm,
	}
}

func (u *httpHandler) Register(c echo.Context) error {
	reqDTO := domain.UserAuthRequest{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	data, err := u.usecase.Register(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (u *httpHandler) Login(c echo.Context) error {
	reqDTO := domain.UserAuthRequest{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	data, err := u.usecase.Login(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	token, err := u.tokenManager.GenerateToken(data.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
