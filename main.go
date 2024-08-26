package main

import (
	"back-end/api"
	"back-end/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal(err)
	}

	file1, err := os.Open("./sql/01-db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	storage.NewDB()

	err = storage.Migrate(file1)
	if err != nil {
		log.Fatal(err)
	}

	api.Run()
}
