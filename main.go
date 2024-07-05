package main

import (
	"api/data"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	router := addRoutes()
	db = initDB()
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}

func initDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	data.SetupDB(db)
	return db
}
