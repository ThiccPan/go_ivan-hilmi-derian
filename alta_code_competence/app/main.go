package main

import (
	"net/http"
	"os"

	"github.com/alta_code_competence/app/config"
	"github.com/alta_code_competence/helper"
	categoryHandler "github.com/alta_code_competence/internal/category/handler"
	categoryRepository "github.com/alta_code_competence/internal/category/repository"
	categoryUsecase "github.com/alta_code_competence/internal/category/usecase"
	itemHandler "github.com/alta_code_competence/internal/item/handler"
	itemRepository "github.com/alta_code_competence/internal/item/repository"
	itemUsecase "github.com/alta_code_competence/internal/item/usecase"
	userHandler "github.com/alta_code_competence/internal/user/handler"
	userRepository "github.com/alta_code_competence/internal/user/repository"
	userUsecase "github.com/alta_code_competence/internal/user/usecase"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbconf := config.DBconf{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}
	db := dbconf.InitDB()
	googleUUID := helper.NewGoogleUUID()
	tokenManager := helper.NewAuthJWT()

	userRepo := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userHandler.NewHttpUserHandler(&userUsecase, tokenManager)

	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUsecase := categoryUsecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := categoryHandler.NewHttpCategoryHandler(&categoryUsecase)

	itemRepo := itemRepository.NewItemRepository(db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepo, googleUUID)
	itemHandler := itemHandler.NewHttpItemHandler(&itemUsecase)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})
	e.GET("/category", categoryHandler.GetAll)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	itemRoute := e.Group("/items")
	itemRoute.Use(echojwt.JWT([]byte(helper.JWT_SECRET)))
	itemRoute.GET("", itemHandler.GetAll)
	itemRoute.GET("/category/:category_id", itemHandler.GetAllByCategory)
	itemRoute.POST("", itemHandler.Create)
	itemRoute.PUT("", itemHandler.Edit)
	itemRoute.GET("/:id", itemHandler.GetById)
	itemRoute.DELETE("/:id", itemHandler.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
