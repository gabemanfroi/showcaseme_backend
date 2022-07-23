package db

import (
	"fmt"
	"github.com/golobby/container/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"showcaseme/domain/models"
	"showcaseme/infra/core"
	"showcaseme/internal/utils"
)

type PostgresSession struct{}

func CreateSqlInstance() *gorm.DB {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		core.AppConfig.DbHost,
		core.AppConfig.DbPort,
		core.AppConfig.DbUser,
		core.AppConfig.DbName,
		core.AppConfig.DbPassword)

	db, err := gorm.Open(postgres.Open(dbUrl))
	utils.Check(err, "Error creating database connection ")

	return db
}

func GetSqlInstance() *gorm.DB {
	var injector *gorm.DB
	utils.Check(container.Resolve(&injector), "Error while retrieving Database instance ")
	return injector
}

func Migrate() {
	log.Printf("Starting Database Migrations...")
	var db *gorm.DB
	err := container.Resolve(&db)

	if err != nil {
		log.Fatal("error while retrieving database instance" + err.Error())
	}

	utils.Check(db.AutoMigrate(&models.User{}), "failed to migrate users")
	utils.Check(db.AutoMigrate(&models.CarouselItem{}), "failed to migrate carousel_item")
	utils.Check(db.AutoMigrate(&models.SkillCategory{}), "failed to migrate skill_categories")
	utils.Check(db.AutoMigrate(&models.Skill{}), "failed to migrate skills")
	log.Printf("Database Migrations Completed...")
}
