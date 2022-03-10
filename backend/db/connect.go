package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"toru/backend/model"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=10.133.0.131 user=postgres password=password dbname=shop port=5432 sslmode=disable"
	data, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	} else {
		fmt.Println("connected database")
		DB = data
		data.AutoMigrate(&model.Product{})
	}
}