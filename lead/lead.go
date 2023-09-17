package lead

import (
	"fiber-crm/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Define the Lead struct with JSON tags
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

// Handle GET request to fetch all leads
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

// Handle GET request to fetch a specific lead by ID
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

// Handle POST request to create a new lead
func CreateLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

// Handle PUT request to update an existing lead by ID
func UpdateLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).Send("No lead found with the given ID")
		return
	}

	updatedLead := new(Lead)
	if err := c.BodyParser(updatedLead); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Update lead fields
	lead.Name = updatedLead.Name
	lead.Company = updatedLead.Company
	lead.Email = updatedLead.Email
	lead.Phone = updatedLead.Phone

	db.Save(&lead)

	c.JSON(lead)
}

// Handle DELETE request to delete a lead by ID
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with the given ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}
