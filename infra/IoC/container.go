package IoC

import (
	"github.com/golobby/container/v3"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"showcaseme/domain/interfaces/services"
	services2 "showcaseme/domain/services"
	"showcaseme/infra/db"
)

type Container struct {
	userService *services.UserService `container:"name"`
}

func InitContainer() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error while trying to read .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")

	err := container.Transient(func() services.UserService { return services2.NewUserService() })
	err = container.Singleton(func() *gorm.DB {
		return db.CreateDB(dbUser, dbPassword, dbPort, dbHost, dbName)
	})
	if err != nil {
		log.Fatal("error while creating the container " + err.Error())
	}
}
