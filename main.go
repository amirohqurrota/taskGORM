package main

import (
	"GORM-2/config"
	"GORM-2/routes"
)

func main() {
	config.InitDB()
	echoApp := routes.NewRoutes()
	echoApp.Start(":8080")

}
