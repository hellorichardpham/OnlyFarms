package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/hellorichardpham/onlyfarms/user-service/controllers"
	"github.com/hellorichardpham/onlyfarms/user-service/models/daos"
)

type application struct {
	users *daos.User
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the home handler")
}

func main() {
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB("root:password@/onlyfarms?parseTime=true")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// app := &application{
	// 	users: &daos.User{DB: db},
	// }

	u := controllers.User{
		UserDAO: &daos.User{DB: db},
	}
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("Get")
	router.HandleFunc("/user/{id}", u.GetUserByID).Methods("Get")
	router.HandleFunc("/user", u.CreateUser).Methods("Post")
	router.HandleFunc("/authenticate", u.AuthenticateUser).Methods("Post")

	err = http.ListenAndServe(":4000", router)
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
