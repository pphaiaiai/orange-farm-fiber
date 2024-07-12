package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Connect to the database
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	if ssl == "" {
		ssl = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, ssl)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Variety{}, &User{}, &Farm{}, &Orange{})

	// Setup Fiber
	app := fiber.New()

	// Middleware
	app.Use("/variety", authRequired)
	app.Use("/user", authRequired)
	app.Use("/farm", authRequired)
	app.Use("/orange", authRequired)

	// Routes
	AuthRoute(app, db)
	VarietyRoute(app, db)
	FarmRoute(app, db)
	OrangeRoute(app, db)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
