package main

import (
	"back-end/api"
	"back-end/storage"
	"time"
)

func main() {
	time.Sleep(time.Second * 30)

	storage.NewDB()

	api.Run()
}
