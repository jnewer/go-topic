package main

import (
	"fmt"
	"gin-gorm-page/config"
	models "gin-gorm-page/model"
	routes "gin-gorm-page/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	// defer db.Close()

	config.DB.AutoMigrate(&models.User{})
	r := routes.SetUpRouter()

	r.Run()
}
