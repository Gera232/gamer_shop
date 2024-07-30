package main

import (
	"back-end/api"
	"back-end/storage"
	"time"
)

func main() {
	time.Sleep(time.Second * 30)

	api.Run()

	storage.NewDB()
}
