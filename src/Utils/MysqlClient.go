package Utils

import (
	"os"

	"github.com/gofiber/storage/mysql"
	"github.com/joho/godotenv"
)

func CreateMysqlClient() *mysql.Storage {
	HOST := "127.0.0.1"
	PORT := 3306
	USER := "topan"
	PASSWORD := "123456"
	DATABASE := "db-klinik"

	// Create a new storage client
	storage := mysql.New(mysql.Config{
		Host:     HOST,
		Port:     PORT,
		Username: USER,
		Password: PASSWORD,
		Database: DATABASE,
	})

	return storage
}

func DatabaseMod() *mysql.Storage {
	//load env file on root folder
	env := godotenv.Load(".env")
	if env != nil {
		panic("Error loading .env file")
	}

	HOST     := os.Getenv("DB_HOST")
	USER     := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	PORT     := 3390
	DATABASE := os.Getenv("DB_DATABASE3")

	// Create a new storage client
	storage := mysql.New(mysql.Config{
		Host:     HOST,
		Port:     PORT,
		Username: USER,
		Password: PASSWORD,
		Database: DATABASE,
	})

	return storage
}


