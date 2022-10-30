package core

import (
	"log"
	"os"
)

type Config struct {
	DbHost               string
	DbPassword           string
	DbName               string
	DbPort               string
	DbUser               string
	AppPort              string
	AwsAccessKeyId       string
	AwsAccessKeyPassword string
	AwsRegion            string
	AwsBucketName        string
}

var AppConfig = Config{}

func LoadConfig() {
	log.Println("Loading Environment Variables...")
	AppConfig.DbHost = os.Getenv("DB_HOST")
	AppConfig.DbPassword = os.Getenv("DB_PASSWORD")
	AppConfig.DbName = os.Getenv("DB_NAME")
	AppConfig.DbPort = os.Getenv("DB_PORT")
	AppConfig.DbUser = os.Getenv("DB_USER")
	AppConfig.AppPort = os.Getenv("APP_PORT")
	AppConfig.AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	AppConfig.AwsAccessKeyPassword = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AppConfig.AwsRegion = os.Getenv("AWS_REGION")
	AppConfig.AwsBucketName = os.Getenv("AWS_BUCKET_NAME")
}
