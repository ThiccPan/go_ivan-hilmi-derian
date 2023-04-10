package test

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

type testCase struct {
	name        string
	path        string
	param       string
	body        string
	expectCode  int
	sizeData    int
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func initTestEnv(req *http.Request) (*httptest.ResponseRecorder, echo.Context) {
	e := InitEchoTestAPI()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}
