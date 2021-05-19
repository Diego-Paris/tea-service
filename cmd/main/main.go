package main

import (
	"fmt"

	"github.com/Diego-Paris/tea-service/pkg/config"
)

func main() {

	fmt.Println(config.Port)
	fmt.Println(config.Environment)
	fmt.Println(config.MongoCredentials)
}
