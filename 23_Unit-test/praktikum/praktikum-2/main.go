package main

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/route"
)

func init() {
	config.InitDB()
}

func main() {
	// create a new echo instance
	e := route.NewInstance()
	// start the server, and log if it fails

	e.Logger.Fatal(e.Start(":8000"))
}
