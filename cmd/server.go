package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/talking_clock/pkg/routes"
	"log"
)

func StartServer() {
	//Routes
	router := gin.New()
	routeGroup := router.Group("/api")
	routes.RegisterRoutes(routeGroup)

	err := router.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
