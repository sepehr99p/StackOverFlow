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

var mockQuestion = models.Question{
	UserId: 123, Description: "some description", Votes: 33, Answers: []*models.Answer{&mockAnswer, &mockAnswer}, TagId: 33,
	QuestionId: 33,
}

var mockAnswer = models.Answer{UserId: 234, Description: "some answer description", Votes: 44}

var mockQuestions = []models.Question{mockQuestion, mockQuestion, mockQuestion}

var database *gorm.DB

func main() {
	//connectDatabase()
	database = connectDatabaseGorm()

	router := gin.Default()
	router.GET("/questions/:id", fetchQuestionById)
	router.GET("/questions/all", fetchQuestions)
	router.GET("/questions/my/:user_id", fetchMyQuestions)
	//router.POST("/questions/delete")
	//router.POST("question/add")
	//router.POST("answer/add")
	//router.POST("answer/delete")
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func fetchQuestionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var questionToFind = models.Question{QuestionId: id}
	result := database.First(&questionToFind)

	result.Error.Error()

	if questionToFind.QuestionId == id {
		c.IndentedJSON(http.StatusOK, questionToFind)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
}

func fetchQuestions(c *gin.Context) {
	var questions []models.Question
	result := database.Find(&questions)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
		return
	}
	c.IndentedJSON(http.StatusOK, questions)
}

func fetchMyQuestions(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var userToFind = models.User{UserId: int(id)}
	var questions []models.Question
	result := database.Model(&userToFind).Find(questions)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
		return
	}
	c.IndentedJSON(http.StatusOK, questions)
}

func connectDatabaseGorm() *gorm.DB {
	sqlDB, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/stackoverflow")
	if err != nil {
		log.Fatalf("error ")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	migratorErr := gormDB.Migrator().AutoMigrate(models.Question{}, models.Answer{}, models.Comment{}, models.User{})
	if migratorErr != nil {
		return nil
	}
	return gormDB
}
