package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/jaysonmulwa/go-rest/internal/transport/http"
)

//Struct contains ponters to database
type App struct {
}

//Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Jayson Worlds")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}
}
