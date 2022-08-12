package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	var err error
	dsn := "테스트:테스트@tcp(127.0.0.1:3306)/테스트?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}
