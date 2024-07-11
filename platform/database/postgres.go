package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pphaiaiai/orange-farm-fiber/app/models"
	"github.com/pphaiaiai/orange-farm-fiber/app/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() {
	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return
	}

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(postgresConnURL), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Variety{})
	fmt.Println("Database migration completed!")
}
