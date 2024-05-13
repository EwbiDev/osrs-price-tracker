package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"time"

	"EwbiDev/osrs-price-tracker/controllers"
	"EwbiDev/osrs-price-tracker/db"
	"EwbiDev/osrs-price-tracker/middleware"

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

	router.Use(middleware.Logger)

	ItemController := controllers.NewItemController(queries, ctx)

	router.HandleFunc("/items/{id:[0-9]+}", ItemController.Get).Methods("GET")
	router.HandleFunc("/items", ItemController.List).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		http.Error(w, "404 Not Found", http.StatusNotFound)
		timeElapsed := time.Since(timeStart)
		log.Printf("%v 404 %v", r.URL.Path, timeElapsed)
	})

	http.Handle("/", router)

	log.Printf("Starting server on http://localhost:4000")
	http.ListenAndServe(":4000", router)
}
