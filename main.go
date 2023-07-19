package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.InitDB()
	config.InitialMigration()
	config.LoadDotEnv()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
