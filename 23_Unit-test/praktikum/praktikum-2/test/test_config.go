package test

import (
	"23_Unit-test/praktikum/praktikum-2/config"

	"github.com/labstack/echo"
)

type testCase struct {
	name       string
	path       string
	param      string
	body       string
	expectCode int
	sizeData   int
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}
