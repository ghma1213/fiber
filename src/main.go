package main

import (
	"fiber/src/database"
	"fiber/src/routes"
	"log"
)

func main() {
	app := routes.Router()
	database.Init()
	log.Fatal(app.Listen(":3000"))
}
