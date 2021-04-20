package server

import (
	"fmt"
	"net/http"

	"github.com/koficypher/simple_server/database"
	transport "github.com/koficypher/simple_server/router"
)

// App struct for declaring our dependencies such as DB connections etc
type App struct{}

// Run method starts up our server
func (app *App) Run() error {
	fmt.Println("Starting up our server.....")

	var err error

	DB, err := database.NewDatabase()

	if err != nil {
		return err
	}

	// err = database.MigrateModels(DB)

	// if err != nil {
	// 	return err
	// }

	// err = database.RunSeeder(DB)

	// if err != nil {
	// 	return err
	// }

	handler := transport.NewHandler()
	handler.SetupRoutes(DB)
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to initialize server")

		return err
	}

	return nil
}
