package main

import (
	"flag"
	"user-management-service/app"
)

func main() {
	migrate := flag.Bool("migrate", false, "Runs database migrations.")
	migrateDown := flag.Bool("migrate-down", false, "Drops the state of the database.")
	flag.Parse()

	a := app.NewApp()

	if *migrate {
		a.MigrateDatabase()
		return
	}
	if *migrateDown {
		a.MigrateDownDatabase()
		return
	}

	a.Run()
}
