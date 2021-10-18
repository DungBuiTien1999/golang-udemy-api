package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/driver"
	"github.com/DungBuiTien1999/udemy-api/internal/handlers"
	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
)

const portNumber = ":9090"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Printf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// connect to database
	log.Println("Connecting to database...")
	connectionString := "root:root@/testudemy?parseTime=true"
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database...")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	helpers.NewHelpers(&app)
	return db, nil
}
