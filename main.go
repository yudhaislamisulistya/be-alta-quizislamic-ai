package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.LoadDotEnv()
	config.InitDB()
	config.InitialMigration()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
