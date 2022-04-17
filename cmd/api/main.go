package main

import (
	"log"
	"rest/server"
)

func main() {
	app := server.NewApp()
	err := app.Run("8080")
	log.Fatal(err)
}
