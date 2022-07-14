package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"showcaseme/application"
	"showcaseme/infra/IoC"
	"showcaseme/infra/core"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

func init() {

	log.Printf("Setting Up Your Server...")

	utils.Check(godotenv.Load(), "Error while trying to read .env file...")

	core.LoadConfig()
	IoC.InitContainer()
	db.Migrate()
	log.Printf("Server Setup complete...")
}

func StartServer() {

	app := fiber.New()

	application.RegisterRoutes(app)

	log.Println(fmt.Sprintf("Starting Server on port %s", core.AppConfig.AppPort))
	log.Fatal(app.Listen(fmt.Sprintf(":%v", core.AppConfig.AppPort)))

}
