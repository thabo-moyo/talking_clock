package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/routes"
	"log"
)

func StartServer() {
	//Routes
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	routeGroup := router.Group("/api")
	routes.RegisterRoutes(routeGroup)

	err := router.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	} else {
		fmt.Sprintln("Server has started")
	}

}
