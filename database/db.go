package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

import (
	"Learning/models"
)

var DB *gorm.DB

func ConnectDatabase() {
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

	// Perform migrations
	err = gormDB.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{}, &models.Comment{}, &models.Tag{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration successful!")
	DB = gormDB
}
