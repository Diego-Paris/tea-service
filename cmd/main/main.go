package main

import (
	"fmt"

	"github.com/Diego-Paris/tea-service/pkg/config"
)

func main() {

	fmt.Printf("%v, %T\n", config.Globals.Port, config.Globals.Port)
	fmt.Println(config.Globals.Environment)
	fmt.Println(config.Globals.MongoCredentials)
}
