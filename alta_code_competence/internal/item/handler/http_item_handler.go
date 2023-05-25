package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alta_code_competence/internal/domain"
	"github.com/alta_code_competence/internal/item/usecase"
	"github.com/labstack/echo/v4"
)

type HttpHandler interface{}

type httpHandler struct {
	usecase usecase.ItemUsecase
}

func NewHttpItemHandler(u usecase.ItemUsecase) *httpHandler {
	return &httpHandler{
		usecase: u,
	}
}

func (h *httpHandler) GetAll(c echo.Context) error {
	queryDTO := c.QueryParam("keyword")
	fmt.Println(queryDTO)
	data, err := h.usecase.GetAll(queryDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (h *httpHandler) GetAllByCategory(c echo.Context) error {
	param, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}
	
	data, err := h.usecase.GetAllByCategory(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (h *httpHandler) Create(c echo.Context) error {
	requestDTO := domain.ItemCreateRequest{}
	err := c.Bind(&requestDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	data, err := h.usecase.Create(requestDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (h *httpHandler) GetById(c echo.Context) error {
	idDTO := c.Param("id")
	data, err := h.usecase.GetById(idDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (h *httpHandler) Edit(c echo.Context) error {
	reqDTO := domain.ItemEditRequest{}
	c.Bind(&reqDTO)
	data, err := h.usecase.Update(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}


func (h *httpHandler) Delete(c echo.Context) error {
	idDTO := c.Param("id")
	data, err := h.usecase.Delete(idDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}