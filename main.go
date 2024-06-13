package main

import (
	"backend/api"
	"backend/storage"
	"log"
	"time"
)

func main() {
	for i := 5; i >= 0; i-- {
		log.Printf("Starting server in %v", i)
		time.Sleep(time.Second * 1)
	}

	storage.NewDB()

	api.Run()
}
