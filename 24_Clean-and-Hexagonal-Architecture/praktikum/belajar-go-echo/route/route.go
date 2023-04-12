package route

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"fmt"

	"github.com/labstack/echo/v4"
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
	app.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
}
