package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_tutorial"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{}, &Photo{})

	DB = db
}
