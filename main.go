package main

import (
	"api/api"
	"api/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./sql/01-db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	storage.NewDB()

	err = storage.Migrate(file)
	if err != nil {
		log.Fatal(err)
	}

	api.Run()
}
