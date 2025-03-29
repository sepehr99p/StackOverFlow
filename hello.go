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

var mockQuestions []models.Question

func main() {
	//connectDatabase()
	connectDatabaseGorm()

	router := gin.Default()
	router.GET("/questions/:id", fetchQuestionById)
	router.GET("/questions/all", fetchQuestions)
	router.GET("/questions/my", fetchMyQuestions)
	//router.POST("/questions/delete")
	//router.POST("question/add")
	//router.POST("answer/add")
	//router.POST("answer/delete")
	router.Run("localhost:8080")
}

func fetchQuestionById(c *gin.Context) {
	id := c.Param("id")
	for _, question := range mockQuestions {
		if id == strconv.FormatInt(question.QuestionId, 10) {
			c.IndentedJSON(http.StatusOK, question)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
}

func fetchQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockQuestions)
}

func fetchMyQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockQuestions)
}

var mockQuestion = models.Question{
	UserId: 123, Description: "some description", Votes: 33, Answers: []models.Answer{mockAnswer, mockAnswer}, TagId: 33,
	QuestionId: 33,
}

var mockAnswer = models.Answer{UserId: 234, Description: "some answer description", Votes: 44}

func connectDatabaseGorm() *gorm.DB {
	//dsn := "root:test@tcp(127.0.0.1:3306)/stackoverflow?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal("failed to connect to gorm")
	//}

	sqlDB, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/stackoverflow")
	if err != nil {
		log.Fatalf("error ")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	result := gormDB.Create(&mockQuestion)

	result.Error.Error()

	return gormDB
	//return db
}

func connectDatabase() {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/stackoverflow")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
