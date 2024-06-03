package main

import (
	"myapp/database"
	"myapp/router"
)

func main() {
	database.ConnectDatabase()
	r := router.SetupRouter()
	r.Run(":8080")
}