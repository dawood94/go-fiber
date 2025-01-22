package main

import (
	"github.com/dawood94/go-fiber/database"
	"github.com/dawood94/go-fiber/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var log = logrus.New() // Fehlerprotokollierung

func setupRoutes(app *fiber.App) {
	app.Get("/api/v2/lead", lead.GetLeads)
	app.Get("/api/v2/lead/:id", lead.GetLead)
	app.Post("/api/v2/lead", lead.NewLead)
	app.Delete("/api/v2/lead/:id", lead.DeleteLead)
}

// to open a connection to the database
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db") // used gorm to open a connection to sqlite3 database ( database name: leads.db)
	if err != nil {
		//panic(("faild to connect database"))
		log.WithError(err).Fatal("Failed to connect to database")
	}
	log.Info("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{}) // Erstellt oder aktualisiert die Tabellen, um sicherzustellen, dass sie den definierten Strukturen(Lead) entsprechen, in der Datenbank vorhanden ist und die richtigen Spalten hat.
	log.Info("Database Migrated")
}

func main() {
	app := fiber.New() // create a new fiber app

	//Fehlerbehandlung in Routen ---> Middkeware wird verwendet , bevor die Routen aufgerufen werden
	app.Use(func(c *fiber.Ctx) error { // Middleware definition

		if err := c.Next(); err != nil {
			log.WithError(err).Error("An Error occuted")
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
			return err

		}
		return nil
	})

	initDatabase() // open a connection to the database
	setupRoutes(app)
	app.Listen(":3000")           // start server on port 3000
	defer database.DBConn.Close() // close database connection when main function ends (defer)

}
