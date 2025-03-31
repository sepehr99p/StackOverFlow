package routes

import (
	"Learning/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Question routes
	router.GET("/questions/:id", handlers.FetchQuestionById)
	router.GET("/questions/all", handlers.FetchQuestions)
	router.POST("/questions/add", handlers.PostQuestion)
	router.GET("/questions/my/:user_id", handlers.FetchMyQuestions)
	router.DELETE("/questions/delete", handlers.DeleteQuestion)

	// User routes
	router.POST("/user/add", handlers.AddUser)
	router.DELETE("/user/delete", handlers.DeleteUser)

	// Answer routes
	router.POST("/answer/add", handlers.AddAnswer)

	//comment routes
	router.POST("/comment/add", handlers.AddComment)

	return router
}
