package main

import (
	"fmt"
	"log"

	"github.com/Diego-Paris/tea-service/pkg/config"
)

func main() {

	err := config.SetupDB()
	if err != nil {
		log.Fatalln("Could not connect to database.\n", err)
	}

	fmt.Println(config.Port)
	fmt.Println(config.Environment)
	fmt.Println(config.MongoCredentials)
}
