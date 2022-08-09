package responose

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{http.StatusOK, message, data})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{http.StatusBadRequest, message, 0})
}
