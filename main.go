package main

import (
	"fmt"
	"github.com/charan678/go-rest-example/config"
	"github.com/charan678/go-rest-example/routes"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	url := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	router.SetupRoutes().Run(url)
}
