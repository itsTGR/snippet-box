package main

import (
	"flag"
	_ "fmt"
	"log/slog"
	"net/http"
	"os"
)

// application structure will hold the dependencies for the web application (dependency injection)
type application struct {
	logger *slog.Logger
}

func main() {

	// Setting input flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Setting logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Instantiating application struct which contains all the web app dependencies
	app := application{
		logger: logger,
	}

	// Setting web server
	mux := app.routes()

	logger.Info("Starting server", "address", *addr)
	err := http.ListenAndServe(*addr, mux)

	logger.Error("Web server failed", slog.String("error", err.Error()))
	os.Exit(1)
}

// Implementing dependency injection if handlers are spread in the packages
// Create an "application" public (exportable) struct that contains all the dependencies for the application
// Make closures that wraps the http handlers, and the outer function will receive the pointer to the configuration, so the http handlers will be able to access all the dependencies.

// package config
// type Application struct {
// 	logger *slog.Logger
//  db *sql.DB
//  metrics *prometheus.Client
// }

// package main
// cfg := &config.Application {
//	logger: logger,
//  db: database,
//  metrics: promClient,
//}

// package XXXX
// func ExampleHandler(app *config.Application) *http.HandlerFunc {
// 		return func (w http.ResponseWriter, r *http.Request) {
//			app.logger.Error(asdfsadf, asdfasdf, asdfasdf)
//  		app.db.Query(fasdfasdfasdfasdf)
//  		app.metrics(asdfasdfasdf)
//		}
//}
