package main

import (
	"25_Docker/praktikum/config"
	"25_Docker/praktikum/route"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	// create a new echo instance
	e := route.NewInstance()
	// start the server, and log if it fails

	e.Logger.Fatal(e.Start(":8000"))
}
