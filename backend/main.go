package main

import (
	"backend/database"
	"backend/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	r.Run(":8080") // buka di http://localhost:8080
}
