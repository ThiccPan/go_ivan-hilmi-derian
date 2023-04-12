package route

import (
	"belajar-go-echo/config"
	constants "belajar-go-echo/constant"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt"
)

func InitRoute(app *echo.Echo) {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("err conn db")
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		fmt.Println("err migr db")
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	usercase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(usercase)

	appJwt := app.Group("")
	appJwt.Use(echojwt.JWT([]byte(constants.JWT_SECRET)))
	appJwt.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
	app.POST("/login", userController.Login)
}
