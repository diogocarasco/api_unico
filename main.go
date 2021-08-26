package main

import (
	configs "api_unico/config"
	"api_unico/database"
)

func main() {

	database.InitializeDB()

	server := configs.GetServer()

	server.Run(":8080")

}
