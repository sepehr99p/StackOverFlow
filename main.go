package main

import (
	"Learning/database"
	_ "Learning/docs"
	routes "Learning/routers"
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
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	} else {
		log.Println("Server running on localhost:8080")
	}
}
