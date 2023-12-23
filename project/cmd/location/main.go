package main

import (
	"github.com/joho/godotenv"
	"log"
	"project/internal/location"
)

func main() {
	a := location_app.NewApp()
	err := godotenv.Load("env.dev")
	if err != nil {
		log.Fatal(err)
	}
	a.Run()
}
