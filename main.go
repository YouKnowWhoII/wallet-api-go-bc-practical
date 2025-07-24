package main

import (
	"log"
	"os"
	"wallet-api-go-bc/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customMiddleware "wallet-api-go-bc/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(customMiddleware.RequestLogger())
	e.Use(customMiddleware.RateLimiter())

	// Register routes
	router.RegisterRoutes(e)

	// Start server
	// Please add PORT environment variable to a .env file, it hasn't been committed for security reasons.
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
