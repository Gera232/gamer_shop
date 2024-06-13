package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./01-database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	storage.NewDB()

	err = storage.Migrate(file)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	api.Run(port)
}
