package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/handler"
)

// RegisterRoutes Register all application routes
func RegisterRoutes(api *gin.RouterGroup) {
	//time routes
	api.GET("/time/human-friendly", handler.GetHumanFriendlyTime)
	api.GET("/time/human-friendly/:time", handler.GetHumanFriendlyTimeByString)
}
