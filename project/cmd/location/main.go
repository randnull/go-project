package main

import (
	"project/internal/location"
)

func main() {
	a := location_app.NewApp()
	//err := godotenv.Load("env.dev")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	a.Run()
}
