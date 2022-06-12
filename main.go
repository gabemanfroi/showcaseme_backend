package main

import (
	"github.com/joho/godotenv"
	"log"
	"showcaseme/infra/IoC"
)

func init() {

	log.Printf("loading settings fron the environment")

	if err := godotenv.Load(); err != nil {
		log.Fatal("error while trying to read .env")
	}
	IoC.InitContainer()
	log.Printf("settings loaded")
}

func main() {

}
