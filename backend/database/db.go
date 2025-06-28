package database

import (
	"fmt"
	"backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/tokped_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Gagal konek DB")
	}
	fmt.Println("✅ Database connected")

	db.AutoMigrate(&models.User{})
	DB = db
}
