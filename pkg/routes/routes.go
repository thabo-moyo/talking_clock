package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/handler"
)

func RegisterRoutes(api *gin.RouterGroup) {
	//time
	api.GET("/time/human-friendly", handler.GetHumanFriendlyTime)
	api.GET("/time/human-friendly/:time", handler.GetHumanFriendlyTimeByString)
}
