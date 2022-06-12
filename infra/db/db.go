package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresSession struct{}

func CreateDB(dbUser, dbPassword, dbPort, dbHost, dbName string) *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open(postgres.Open(dbUrl))

	if err != nil {
		log.Fatal("error while trying to open database connection" + err.Error())
	}

	return db
}
