package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.InitDB()
	config.InitialMigration()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
