package main

import (
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/runner"
	"ksp.sk/proboj/web/web"
)

func main() {
	err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	err = database.OpenDatabase()
	if err != nil {
		panic(err)
	}

	go runner.GeneratorLoop()

	web.Start()
}
