package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Port that the server will be listening on.
	Port int

	// Server Environment, development or production.
	Environment string

	// URI Link to connect with MongoDB.
	MongoCredentials string
)

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
	Port = port

	// Set Server Environment
	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv != "development" && serverEnv != "production" {
		log.Fatal("Invalid server environment, must be development or production")
	}
	Environment = serverEnv

	// Set MongoDB URI
	MongoURI := os.Getenv("MONGO_CREDENTIALS")
	if len(MongoURI) <= 0 {
		log.Fatal("Invalid Mongo URI, must not be empty")
	}
	MongoCredentials = MongoURI

}
