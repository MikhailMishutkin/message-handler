package main

import (
	"github.com/joho/godotenv"
	"log"
	"message_handler/configs"
	"message_handler/internal/app"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.StartService(*config); err != nil {
		log.Fatal(err)
	}
}
