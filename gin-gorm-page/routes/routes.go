package routes

import (
	Controller "gin-gorm-page/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	u1 := r.Group("/admin")
	{
		u1.GET("all", Controller.GetAllUsers)
	}

	return r
}
