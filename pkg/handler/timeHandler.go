package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/utils"
	"net/http"
)

func GetHumanFriendlyTime(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": utils.GetCurrentTime()})
}

func GetHumanFriendlyTimeByString(context *gin.Context) {
	timeByString, err := utils.GetTimeByString(context.Param("time"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect time format"})
		context.Abort()
	} else {
		context.JSON(http.StatusOK, gin.H{"data": timeByString})
	}
}
