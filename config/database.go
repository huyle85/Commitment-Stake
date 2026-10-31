package config

import (
	"fmt"
	"log"
	"os"

	"commitment-stake/models" 

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Lỗi tải file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối Database! Lỗi: ", err)
	}


	err = database.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{})
	if err != nil {
		log.Fatal("Lỗi khi tạo bảng: ", err)
	}

	DB = database
	fmt.Println("Kết nối & Migrate Database thành công!")
}