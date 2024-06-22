package main

import (
	"back-end/api"
	"back-end/storage"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	storage.NewDB()

	api.Run()
}
