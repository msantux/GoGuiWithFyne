package main

import (
	"GoGuiWithFyne/Section_05_GoldWatcher_Part_02/repository"
	"database/sql"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	Toolbar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HttpClient          *http.Client
}

func main() {
	var myApp Config

	// create a fyne application
	fyneApp := app.NewWithID("com.santux.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HttpClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open connection to the database
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a database repository
	myApp.setupDB(sqlDB)

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()
	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
	}
	app.InfoLog.Println("db in:", path)

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(db *sql.DB) {
	app.DB = repository.NewSQLiteRepository(db)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
