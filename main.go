package main

import (
	"api-account/api"
	"api-account/storage"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.Run()
}
