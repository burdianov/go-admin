package database

import (
	"github.com/burdianov/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not establish a database connection")
	}

	DB = db

	db.AutoMigrate(&models.User{})
}
