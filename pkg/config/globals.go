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
		log.Fatalln("Error loading .env file\n", err)
	}

	// Set Port
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalln("Invalid port:\n", err)
	}
	Port = port

	// Set Server Environment
	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv != "development" && serverEnv != "production" {
		log.Fatalln("Invalid server environment, must be development or production")
	}
	Environment = serverEnv

	// Set MongoDB URI
	MongoURI := os.Getenv("MONGO_CREDENTIALS")
	if len(MongoURI) <= 0 {
		log.Fatalln("Invalid Mongo URI, must not be empty")
	}
	MongoCredentials = MongoURI

}
