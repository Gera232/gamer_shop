package main

import (
	"backend/api"
	"backend/storage"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.Run()
}
