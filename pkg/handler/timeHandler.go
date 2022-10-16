package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/utils"
	"net/http"
)

// GetHumanFriendlyTime Returns human friendly time JSON response
func GetHumanFriendlyTime(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": utils.GetCurrentTime()})
}

// GetHumanFriendlyTimeByString Returns given human friendly time JSON response
func GetHumanFriendlyTimeByString(context *gin.Context) {
	timeByString, err := utils.GetTimeByString(context.Param("time"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect time format"})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": timeByString})
	}
}
