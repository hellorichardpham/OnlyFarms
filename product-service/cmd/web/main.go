package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hellorichardpham/onlyfarms/product-service/pkg/models/mysql"
	"github.com/hellorichardpham/onlyfarms/product-service/utilities"
)

func showHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Only Farms Home"))
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	items    *mysql.ItemModel
	packages *mysql.PackageModel
	farmers  *mysql.FarmerModel
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am alive in the produce-service.")
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB("root:password@/onlyfarms?parseTime=true")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		items:    &mysql.ItemModel{DB: db},
		packages: &mysql.PackageModel{DB: db},
		farmers:  &mysql.FarmerModel{DB: db},
	}

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(showHomepage))
	mux.Get("/item/:id", http.HandlerFunc(app.getItem))
	mux.Get("/farmer/:id", http.HandlerFunc(app.getFarmer))
	mux.Get("/items/package/:id", http.HandlerFunc(app.getPackageItems))
	mux.Get("/package/:id", http.HandlerFunc(app.getPackage))
	mux.Get("/healthcheck", http.HandlerFunc(healthCheck))
	mux.Get("/configuration", http.HandlerFunc(utilities.Configuration))

	utilities.RegisterServiceWithConsul()
	err = http.ListenAndServe(utilities.Port(), mux)
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
