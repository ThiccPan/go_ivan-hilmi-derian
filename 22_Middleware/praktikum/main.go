package main

import (
	"22_Middleware/praktikum/config"
	"22_Middleware/praktikum/route"
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
