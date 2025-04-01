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
		ready <- true // Signal readiness
	}()
	<-ready

	// Setup and start server
	router := routes.SetupRouter()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(2)))

	log.Println("Server running on localhost:8080")
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

//  curl http://localhost:8080/questions/all

//curl http://localhost:8080/questions/add \
//--include \
//--header "Content-Type: application/json" \
//--request "POST" \
//--data '{"question_id": 4,"user_id": 33,"tag_idz": 22,"description": "des","votes":33}'

//curl http://localhost:8080/user/add \
//--include \
//--header "Content-Type: application/json" \
//--request "POST" \
//--data '{"user_name": "sepehr","user_id": 33,"reputation": 22}'

//curl http://localhost:8080/answer/add \
//--include \
//--header "Content-Type: application/json" \
//--request "POST" \
//--data '{"answer_id": 31,"question_id": 1,"user_id": 33,"description":"some answer","votes":1}'
