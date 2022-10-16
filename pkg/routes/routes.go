package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	//Projects
	api.GET("/time")
	api.GET("/time/{input}")
}
