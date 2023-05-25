package handler

import (
	"net/http"

	"github.com/alta_code_competence/internal/category/usecase"
	"github.com/labstack/echo/v4"
)

type HttpHandler interface {}

type httpHandler struct {
	usecase usecase.CategoryUsecase
}

func NewHttpCategoryHandler(u usecase.CategoryUsecase) *httpHandler {
	return &httpHandler{
		usecase: u,
	}
}

func (h httpHandler) GetAll(c echo.Context) error {
	data, err := h.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"data": data,
	})
}