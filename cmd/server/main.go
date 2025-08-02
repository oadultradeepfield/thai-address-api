package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/oadultradeepfield/thai-address-api/internal/routes"
)

func main() {
	// Setup GORM with SQLite
	db, err := gorm.Open(sqlite.Open("thai_address.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Setup Echo
	e := echo.New()

	// Setup routes with db
	routes.BaseRoutes(e, db)

	// Start server
	log.Println("Server started on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
