package main

import (
	"belajar-go-echo/route"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	route.InitRoute(app)
	app.Logger.Fatal(app.Start(":8000"))
}
