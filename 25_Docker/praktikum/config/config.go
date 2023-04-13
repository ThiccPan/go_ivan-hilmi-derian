package config

import (
	"25_Docker/praktikum/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		"db",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		"5432",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Post{})
}
