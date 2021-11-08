package main

import "fmt"

//Struct contains ponters to database
type App struct {
}

//Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
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
