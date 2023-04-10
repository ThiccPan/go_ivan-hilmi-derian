package util

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func InitValidateConf(e *echo.Echo)  {
	e.Validator = &CustomValidator{validator: validator.New()}
}