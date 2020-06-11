package app

import (
	"fmt"
	"log"
	"net/http"
	"user-management-service/api"
	"user-management-service/config"
	"user-management-service/database"
	"user-management-service/store"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	config config.Config
	db     database.Postgres
}

func NewApp() App {
	app := App{}

	cfg := config.NewConfiguration("config/config.json")

	cfg.OverrideFromEnvironment()

	app.config = cfg
	app.db = database.NewPostgres(app.config.Database.DbHost, app.config.Database.DbPort,
		app.config.Database.DbUser, app.config.Database.DbPass, app.config.Database.DbName)

	store := store.NewStore(app.db)

	server := api.NewServer(store)

	app.router = api.NewRouter(server)

	return app
}

func (app *App) Run() {
	log.Println("Starting http server on port: ", app.config.HttpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app.config.HttpPort), app.router))
}

// Setup database tables
func (app *App) MigrateDatabase() {
	log.Println("Migrating database ...")
	app.db.CreateGroupTable()
	app.db.CreateUserTable()
}

// Drop database tables
func (app *App) MigrateDownDatabase() {
	log.Println("Dropping tables ...")
	app.db.DropUserTable()
	app.db.DropGroupTable()
}
