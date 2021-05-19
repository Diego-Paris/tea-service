package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type environmentVariables struct {

	// Port that the server will be listening on
	Port int

	// Development or production
	Environment string

	// URI Link to connect with MongoDB
	MongoCredentials string
}

var Globals environmentVariables

func init() {

	// Load in environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Set Port
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("Invalid port: ", err)
	}
	Globals.Port = port

	// Set Server Environment
	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv != "development" && serverEnv != "production" {
		log.Fatal("Invalid server environment, must be development or production")
	}
	Globals.Environment = serverEnv

	// Set MongoDB URI
	MongoURI := os.Getenv("MONGO_CREDENTIALS")
	if len(MongoURI) <= 0 {
		log.Fatal("Invalid Mongo URI, must not be empty")
	}
	Globals.MongoCredentials = MongoURI

}
