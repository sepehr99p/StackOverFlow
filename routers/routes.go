package routes

import (
	"Learning/handlers"
	"Learning/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	authRoute := router.Group("/auth")
	authRoute.POST("/register", handlers.RegisterHandler)
	authRoute.POST("/login", handlers.LoginHandler)
	authRoute.GET("/protected", handlers.ProtectedHandler)

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// Question routes
		protected.GET("/questions/:id", handlers.FetchQuestionById)
		protected.GET("/questions/all", handlers.FetchQuestions)
		protected.POST("/questions/add", handlers.PostQuestion)
		protected.GET("/questions/my", handlers.FetchMyQuestions)
		protected.GET("/questions/voteUp/:id}", handlers.VoteUpQuestion)
		protected.GET("/questions/voteDown/:id}", handlers.VoteDownQuestion)

		// Answer routes
		protected.POST("/answer/add", handlers.AddAnswer)
		protected.GET("/answer/correctAnswer/:id", handlers.CorrectAnswer)
		protected.GET("/answer/voteUp/:id}", handlers.VoteUpAnswer)
		protected.GET("/answer/voteDown/:id}", handlers.VoteDownAnswer)
		protected.GET("/answer/delete", handlers.DeleteAnswer)

		//comment routes
		protected.POST("/comment/add", handlers.AddComment)
		protected.DELETE("/comment/delete", handlers.DeleteComment)

		//tag routes
		protected.POST("/tag/add", handlers.AddTag)
		protected.POST("/tag/questions/all", handlers.FetchTagQuestions)
	}

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware())
	{
		adminRoutes.DELETE("/questions/delete", handlers.DeleteQuestion)

		// User routes
		adminRoutes.POST("/user/add", handlers.AddUser)
		adminRoutes.DELETE("/user/delete", handlers.DeleteUser)
	}

	return router
}
