package config

import "gorm.io/gorm"

var DB *gorm.DB
var DSN string = "root:@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"