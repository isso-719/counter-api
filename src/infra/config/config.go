package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func init() {
	loadDotEnv()
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
	}
}

type HTTPConfig struct {
	Host string
	Port string
}

func LoadHTTPConfig() *HTTPConfig {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		panic("Cannot find ENV: HOST")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		panic("Cannot find ENV: PORT")
	}

	return &HTTPConfig{
		Host: host,
		Port: port,
	}
}

type DatastoreConfig struct {
	ProjectID string
}

func LoadDatastoreConfig() *DatastoreConfig {
	projectID, ok := os.LookupEnv("GOOGLE_PROJECT_ID")
	if !ok {
		panic("Cannot find ENV: PROJECT_ID")
	}
	return &DatastoreConfig{
		ProjectID: projectID,
	}
}
