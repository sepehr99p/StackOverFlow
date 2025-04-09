package database

import (
	"Learning/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

func FetchQuestionsWithAnswersAndComments(questions []models.Question) []gin.H {
	var questionResponses []gin.H
	for _, question := range questions {
		questionResponse := gin.H{"question": question}
		questionResponse["answers"] = FetchAnswersForQuestion(strconv.FormatInt(question.QuestionId, 10))
		var comments []models.Comment
		DB.Where("parent_id = ? AND parent_type = ?", question.QuestionId, "question").Find(&comments)
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
		DB.Where("parent_id = ? AND parent_type = ?", answer.AnswerId, "answer").Find(&answerComments)

		answerResponse := gin.H{
			"answer":   answer,
			"comments": answerComments,
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

func VoteUpAnswerWithOwner(answer *models.Answer) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var answerOwner models.User
		if err := tx.First(&answerOwner, answer.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if err := tx.Model(&answerOwner).Update("reputation", gorm.Expr("reputation + ?", 10)).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&answer).Update("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to vote up: %w", err)
		}

		//todo log the action later
		return nil
	})
	return err
}

func VoteDownAnswerWithOwner(answer *models.Answer) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var answerOwner models.User
		if err := tx.First(&answerOwner, answer.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if answerOwner.Reputation <= 0 {
			return fmt.Errorf("answer owner cannot have negative reputation")
		}

		newReputation := answerOwner.Reputation - 10
		if newReputation < 0 {
			newReputation = 0
		}

		if err := tx.Model(&answerOwner).Update("reputation", newReputation).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&answer).Update("votes", gorm.Expr("votes - ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to downvote: %w", err)
		}

		return nil
	})
	return err
}
