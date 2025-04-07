package routes

import (
	"Learning/handlers"
	"Learning/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/protected", handlers.ProtectedHandler)

	adminRoutes := router.Group("/admin")

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// Question routes
		router.GET("/questions/:id", handlers.FetchQuestionById)
		router.GET("/questions/all", handlers.FetchQuestions)
		router.POST("/questions/add", handlers.PostQuestion)
		router.GET("/questions/my/:user_id", handlers.FetchMyQuestions)
		router.GET("/questions/voteUp/:id}", handlers.VoteUpQuestion)

		// Answer routes
		router.POST("/answer/add", handlers.AddAnswer)
		router.GET("/answer/correctAnswer/:id", handlers.CorrectAnswer)
		router.GET("/answer/voteUp/:id}", handlers.VoteUpAnswer)
		router.GET("/answer/delete", handlers.DeleteAnswer)

		//comment routes
		router.POST("/comment/add", handlers.AddComment)
		router.DELETE("/comment/delete", handlers.DeleteComment)

		//tag routes
		router.POST("/tag/add", handlers.AddTag)
		router.POST("/tag/questions/all", handlers.FetchTagQuestions)
	}

	adminRoutes.Use(middleware.AdminMiddleware())
	{
		router.DELETE("/questions/delete", handlers.DeleteQuestion)

		// User routes
		router.POST("/user/add", handlers.AddUser)
		router.DELETE("/user/delete", handlers.DeleteUser)
	}

	return router
}
