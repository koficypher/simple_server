package main

import (
	"fmt"
	"net/http"

	transport "github.com/koficypher/simple_server/router"
)

// App struct for declaring our dependencies such as DB connections etc
type App struct{}

// Run method starts up our server
func (app *App) Run() error {
	fmt.Println("Starting up our server.....")

	handler := transport.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to initialize server")

		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up server....")
		fmt.Println(err)
	}

}
