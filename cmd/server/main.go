package main

import "fmt"

// App struct for declaring our dependencies such as DB connections etc
type App struct{}

// Run method starts up our server
func (app *App) Run() error {
	fmt.Println("Starting up our server.....")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up server....")
		fmt.Println(err)
	}

}
