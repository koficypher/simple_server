package main

import (
	"fmt"

	"github.com/koficypher/simple_server/kernel/server"
)

func main() {
	app := server.App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up server....")
		fmt.Println(err)
	}
}
