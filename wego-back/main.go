package main

import (
	database "wego-back/database"
	routes "wego-back/routes/routes"
)

func main() {
	database.InitDatabase()
	routes.HandleRquest()
}
