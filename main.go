package main

import (
	"fiber-crm/database"
	"fiber-crm/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	// Define API routes for managing leads
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.CreateLead)
	app.Put("/api/v1/leads/:id", lead.UpdateLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}

func initDatabase() {
	// Initialize and migrate the database
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to the database")
	}
	println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	println("Database migrations applied")
}

func main() {
	// Create a new Fiber application
	app := fiber.New()

	// Initialize the database connection
	initDatabase()

	// Set up API routes
	setupRoutes(app)

	// Start the application on port 3505
	err := app.Listen(3505)
	if err != nil {
		panic("failed to start the application")
	}
	defer database.DBConn.Close()
}
