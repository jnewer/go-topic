package dao

import (
	"gin-gorm-vuejs/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.User)
	ListUser() []models.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}

func (mgr *manager) AddUser(user *models.User) {
	mgr.db.Create(user)
}

func (mgr *manager) ListUser() []models.User {
	var users []models.User
	mgr.db.Find(&users)
	return users
}
