package routers

import (
	"gin-gorm-vuejs/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Start() {
	e := gin.Default()
	newCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},

		MaxAge: 24 * time.Hour,
	})
	e.Use(newCORS)

	e.GET("/users", api.ListUser)
	e.POST("add", api.AddUser)
	e.Run()
}
