package config

import (
	"23_Unit-test/praktikum/praktikum-2/model"
	"fmt"

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

	config := Config{
		DB_Username: "thiccpan",
		DB_Password: "dbpsqlpass432",
		DB_Port:     "5432",
		DB_Host:     "localhost",
		DB_Name:     "praktikum_16",
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DB_Host, config.DB_Username, config.DB_Password, config.DB_Name, config.DB_Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Post{})
}

func InitDBTest() {
	config := Config{
		DB_Username: "thiccpan",
		DB_Password: "dbpsqlpass432",
		DB_Port:     "5432",
		DB_Host:     "localhost",
		DB_Name:     "praktikum_23_test",
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DB_Host, config.DB_Username, config.DB_Password, config.DB_Name, config.DB_Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&model.User{}, &model.Book{}, &model.Post{})
	DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Post{})
}