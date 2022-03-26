package main

import "fmt"

// App - the struct which contains thinks like
// pointers to database connections
type App struct{}

// Run - sets up our application
funct (app *App) Run() error {
	fmt.Println("Setting up our app")
	return nil
}

func main() {
	fmt.Println("Go REST API ")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
