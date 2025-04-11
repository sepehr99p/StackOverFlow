package routes

import (
	"Learning/handlers"
	"Learning/handlers/answer_handler"
	"Learning/handlers/question_handler"
	"Learning/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		protected.GET("/questions/:id", question_handler.FetchQuestionById)
		protected.GET("/questions/all", question_handler.FetchQuestions)
		protected.POST("/questions/add", question_handler.PostQuestion)
		protected.GET("/questions/my", question_handler.FetchMyQuestions)
		protected.GET("/questions/voteUp/:id", question_handler.VoteUpQuestion)
		protected.GET("/questions/voteDown/:id", question_handler.VoteDownQuestion)
		protected.GET("/questions/search", question_handler.SearchQuestions)

		// Answer routes
		protected.POST("/answer_handler/add", answer_handler.AddAnswer)
		protected.GET("/answer_handler/correctAnswer/:id", answer_handler.CorrectAnswer)
		protected.GET("/answer_handler/voteUp/:id", answer_handler.VoteUpAnswer)
		protected.GET("/answer_handler/voteDown/:id", answer_handler.VoteDownAnswer)
		protected.DELETE("/answer_handler/delete", answer_handler.DeleteAnswer)

		//comment routes
		protected.POST("/comment/add", handlers.AddComment)
		protected.DELETE("/comment/delete", handlers.DeleteComment)

		//tag routes
		protected.POST("/tag/add", handlers.AddTag)
		protected.POST("/tag/questions/all", handlers.FetchTagQuestions)

		//report routes
		protected.POST("/report", handlers.ReportInteraction)
	}

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware())
	{
		adminRoutes.DELETE("/questions/delete", question_handler.DeleteQuestion)

		// User routes
		adminRoutes.POST("/user/add", handlers.AddUser)
		adminRoutes.DELETE("/user/delete", handlers.DeleteUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(2)))

	return router
}
