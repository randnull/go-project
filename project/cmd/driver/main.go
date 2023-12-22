package main

import (
	"github.com/joho/godotenv"
	"log"
	"project/internal/driver"
)

func main() {
	a := driver_app.NewApp()
	err := godotenv.Load("env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	a.Run()
}
