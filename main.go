package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"

	"EwbiDev/osrs-price-tracker/controllers"
	"EwbiDev/osrs-price-tracker/db"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()
	ctx := context.Background()
	dbInit, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	queries := db.New(dbInit)

	ItemController := controllers.NewItemController(queries, ctx)

	router.HandleFunc("/items/{id:[0-9]+}", ItemController.Get).Methods("GET")

	http.Handle("/", router)

	log.Printf("Starting server on http://localhost:4000")
	http.ListenAndServe(":4000", router)
}
