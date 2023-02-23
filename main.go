package main

import (
	"hearxtest/api"
	"hearxtest/db"
	"hearxtest/env"
)

func main() {
	env.LoadEnv()
	db := db.InitPostgresDB()
	app := api.InitAPI(db)
	app.Run(":5050")
}
