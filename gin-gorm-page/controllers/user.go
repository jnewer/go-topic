package controllers

import (
	dao "gin-gorm-page/dao"
	models "gin-gorm-page/model"
	utils "gin-gorm-page/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.User
	userLists, err := dao.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})
}
