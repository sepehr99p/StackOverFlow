package db_helper

import (
	"Learning/database"
	"Learning/models"
	"fmt"
	"gorm.io/gorm"
)

func VoteDownQuestion(question *models.Question) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var questionOwner models.User
		if err := tx.First(&questionOwner, question.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if questionOwner.Reputation <= 0 {
			return fmt.Errorf("answer owner cannot have negative reputation")
		}

		newReputation := questionOwner.Reputation - 10
		if newReputation < 0 {
			newReputation = 0
		}

		if err := tx.Model(&questionOwner).Update("reputation", newReputation).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&question).Update("votes", gorm.Expr("votes - ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to downvote: %w", err)
		}

		return nil
	})
	return err
}

func VoteUpQuestion(question *models.Question) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var questionOwner models.User
		if err := tx.First(&question, question.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if err := tx.Model(&questionOwner).Update("reputation", gorm.Expr("reputation + ?", 10)).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&question).Update("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to vote up: %w", err)
		}

		//todo log the action later
		return nil
	})
	return err
}

func VoteUpAnswerWithOwner(answer *models.Answer) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
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
	err := database.DB.Transaction(func(tx *gorm.DB) error {
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
