package database

import (
	"Learning/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func FetchQuestionsWithAnswersAndComments(questions []models.Question) []gin.H {
	var questionResponses []gin.H
	for _, question := range questions {
		questionResponse := gin.H{"question_handler": question}
		questionResponse["answers"] = FetchAnswersForQuestion(strconv.FormatInt(question.QuestionId, 10))
		var comments []models.Comment
		DB.Where("parent_id = ? AND parent_type = ?", question.QuestionId, "question_handler").Find(&comments)
		questionResponse["comments"] = comments
		questionResponses = append(questionResponses, questionResponse)
	}
	return questionResponses
}

func FetchAnswersForQuestion(questionId string) []gin.H {
	var answers []models.Answer
	DB.Where("question_id = ?", questionId).Find(&answers)
	var answersWithComments []gin.H
	for _, answer := range answers {
		var answerComments []models.Comment
		DB.Where("parent_id = ? AND parent_type = ?", answer.AnswerId, "answer_handler").Find(&answerComments)

		answerResponse := gin.H{
			"answer_handler": answer,
			"comments":       answerComments,
		}
		answersWithComments = append(answersWithComments, answerResponse)
	}
	return answersWithComments
}

func IsUserAlreadyExist(phoneNumber string) bool {
	var count int64
	DB.Model(&models.User{}).
		Where("user_name = ?", phoneNumber).
		Count(&count)
	return count > 0
}

func CreateUser(user models.UserRegister) *string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		message := "Error hashing password"
		return &message
	}
	newUser := models.User{
		UserName: user.PhoneNumber,
		Password: string(hashedPassword),
	}
	result := DB.Create(&newUser)
	if result.Error != nil {
		message := result.Error.Error()
		return &message
	}
	return nil
}
