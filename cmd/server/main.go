package main

import (
	"fmt"
	"net/http"

	"github.com/lil-newty/go-rest-api-01/internal/comment"
	"github.com/lil-newty/go-rest-api-01/internal/database"
	transportHTTP "github.com/lil-newty/go-rest-api-01/internal/transport/http"
)

// App - the struct which contains thinks like
// pointers to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err = database.MigrateDB(db); err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":9090", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API ")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up Our REST API")
		fmt.Println(err)
	}
}
