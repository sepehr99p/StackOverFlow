package main

import (
	"Learning/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var database *gorm.DB

func main() {
	database = connectDatabaseGorm()
	if database == nil {
		log.Fatal("Database connection is nil! Exiting...")
	}

	router := gin.Default()
	router.GET("/questions/:id", fetchQuestionById)
	router.GET("/questions/all", fetchQuestions)
	router.GET("/questions/my/:user_id", fetchMyQuestions)
	router.POST("/questions/add", postQuestion)
	router.POST("/user/add", addUser)

	log.Println("Server running on localhost:8080")
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	//router.POST("/questions/delete")

	//router.POST("answer/add",)
	//router.POST("answer/delete")
}

func postQuestion(c *gin.Context) {
	var question models.Question

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	// 🔍 Check if user exists
	var user models.User
	if err := database.First(&user, question.UserId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User does not exist", "error": err.Error()})
		return
	}

	// Insert question
	result := database.Create(&question)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating question", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, question)
}

func addUser(c *gin.Context) {
	var user models.User
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	// Insert question
	result := database.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func fetchQuestionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var questionToFind = models.Question{QuestionId: id}
	result := database.First(&questionToFind)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error on fetching data from database"})
		return
	}

	if questionToFind.QuestionId == id {
		c.IndentedJSON(http.StatusOK, questionToFind)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
}

func fetchQuestions(c *gin.Context) {
	var questions []models.Question
	result := database.Find(&questions)

	// Log the error to see what's causing the 500 error
	if result.Error != nil {
		log.Println("Error fetching questions:", result.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving questions"})
		return
	}

	// Check if no questions were found
	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}

func fetchMyQuestions(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var userToFind = models.User{UserId: int(id)}
	var questions []models.Question

	// FIX: Pass &questions instead of questions
	result := database.Model(&userToFind).Where("user_id = ?", id).Find(&questions)

	if result.Error != nil {
		log.Println("Error fetching user questions:", result.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user's questions"})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found for this user"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}

func connectDatabaseGorm() *gorm.DB {
	sqlDB, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/stackoverflow")
	if err != nil {
		log.Fatalf("Failed to open SQL connection: %v", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// 🔥 MIGRATE TABLES IN THE CORRECT ORDER 🔥
	err = gormDB.AutoMigrate(&models.User{}) // User first
	if err != nil {
		log.Fatalf("Migration failed (User): %v", err)
	}

	err = gormDB.AutoMigrate(&models.Question{}) // Then Question
	if err != nil {
		log.Fatalf("Migration failed (Question): %v", err)
	}

	err = gormDB.AutoMigrate(&models.Answer{}) // Then other tables
	if err != nil {
		log.Fatalf("Migration failed (Other tables): %v", err)
	}

	log.Println("Database migration successful!")
	return gormDB
}

//curl http://localhost:8080/question/add \
//--include \
//--header "Content-Type: application/json" \
//--request "POST" \
//--data '{"question_id": 4,"user_id": 33,"tag_idz": 22,"description": "des","votes":33}'

//curl http://localhost:8080/question/add \
//--include \
//--header "Content-Type: application/json" \
//--request "POST" \
//--data '{"user_name": "sepehr","user_id": 33,"reputation": 22}'
