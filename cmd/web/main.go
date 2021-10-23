package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/driver"
	"github.com/DungBuiTien1999/udemy-api/internal/handlers"
	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
	"github.com/DungBuiTien1999/udemy-api/internal/validator"
	_ "github.com/joho/godotenv/autoload"
)

// const portNumber = ":9090"

var PORT = os.Getenv("PORT")

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Printf("Starting application on port %s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	// read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "todoapp", "Database name")
	dbUser := flag.String("dbuser", "root", "Database user")
	dbPass := flag.String("dbpass", "root", "Database password")
	dbPort := flag.String("dbport", "3306", "Database port")
	dbSSL := flag.String("dbssl", "skip-verify", "Database ssl settings (skip-verify, preferred)")

	flag.Parse()
	if *dbName == "" || *dbUser == "" || *dbPass == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	app.InProduction = *inProduction

	app.Validator = validator.NewValidation()

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s", *dbUser, *dbPass, *dbHost, *dbPort, *dbName, *dbSSL)
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
