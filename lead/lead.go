package lead

import (
	"github.com/dawood94/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model

	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error { // nachdem man das Rout aufgerufen hat , muss man das Context Ctx in Fiber übergeben .Dh c hat allr Informationen über die Request, die von User kommt.
	db := database.DBConn
	var leads []Lead
	db.Find(&leads) // Find all leads in the database and store them in the leads variable
	err := db.Find(&leads).Error
	if err != nil {
		return c.Status(500).SendString("Error by getting leads") // If there is an error, send a status code of 500 and the error message
	}
	return c.JSON(leads) // Response in JSON format

}

func GetLead(c *fiber.Ctx) error {

	id := c.Params("id")  // gibt den Wert des Parameters "id" zurück, der in der URL angefordert wurde.
	db := database.DBConn // gibt die Datenbankverbindung zurück, die in database/database.go definiert ist.
	var lead Lead
	db.Find(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with giver ID")

	}
	return c.JSON(lead) // Response in JSON format

}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)                          // create a new Lead struct
	if err := c.BodyParser(lead); err != nil { // Parse the request body of the data from the user (Postman) . because the data is new and  is in JSON format, we need to parse it to a struct
		return c.Status(503).SendString(err.Error()) // If there is an error, send a status code of 503 and the error message

	}
	db.Create(&lead)
	return c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id) // Find the first lead in the database with the id that was requested
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with given ID")

	}
	db.Delete(&lead)
	return c.SendString("Lead Successfully deleted")

}
