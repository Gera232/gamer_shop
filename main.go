package main

import (
	"back-end/api"
	"back-end/storage"
)

func main() {
	storage.NewDB()

	api.Run()
}
