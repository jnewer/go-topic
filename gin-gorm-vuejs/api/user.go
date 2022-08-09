package api

import (
	"gin-gorm-vuejs/dao"
	"gin-gorm-vuejs/models"
	"gin-gorm-vuejs/responose"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		responose.Failed("参数错误", c)
		return
	}

	dao.Mgr.AddUser(&user)

	responose.Success("添加成功", user, c)
}

func ListUser(c *gin.Context) {
	users := dao.Mgr.ListUser()
	responose.Success("查询成功", users, c)
}
