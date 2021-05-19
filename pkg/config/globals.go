package config

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Port that the server will be listening on.
	Port int

	// Environment, development or production.
	Environment string

	// MongoCredentials is the URI Link to connect with MongoDB.
	MongoCredentials string
)

func init() {

	// Load in environment variables
	// Regex allows to run through standard and debug consoles
	// TODO check if this adds significant complexity and if it helps debug servers
	projectDirName := "tea-service"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
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
