package main

import (
	"github.com/Led-New/Gin-Go/data-base"
	//"github.com/Led-New/Gin-Go/models"
	"github.com/Led-New/Gin-Go/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequest()
}
