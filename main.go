package main

import (
	"Learning/database"
	_ "Learning/docs"
	routes "Learning/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	ready := make(chan bool)
	go func() {
		database.ConnectDatabase()
		ready <- true
	}()
	<-ready

	router := routes.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(2)))

	log.Println("Server running on localhost:8080")
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
