package database

import (
	"fmt"
	"log"
	"os"
	"time"

	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"api_unico/models"
)

var DB *gorm.DB

// InitializeDB create a new connection to the database
func InitializeDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Printf("Database connection error: %s \n", err)
	} else {
		fmt.Print("Database connected!\n")
	}

	db.AutoMigrate(&models.Feiras{})

	DB = db

}
