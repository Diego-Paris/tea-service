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
	Port string

	// Environment development or production.
	Environment string

	// MongoCredentials is the URI Link to connect with MongoDB.
	MongoCredentials string

	// DBName is the name of the database that contains our collections
	DBName string
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
	_, err = strconv.Atoi(portString) // check if number
	if err != nil {
		log.Fatalln("Invalid port:\n", err)
	}
	Port = ":" + portString

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

	// Set name of database
	databaseString := os.Getenv("DB_NAME")
	if len(databaseString) <= 0 {
		log.Fatalln("Invalid database name, must not be empty")
	}
	DBName = databaseString
}
