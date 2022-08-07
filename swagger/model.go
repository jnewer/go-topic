package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" example:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" example:"desc" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status" example:"Active"`
}

type Response struct {
	Msg  string
	Data interface{}
}
