package main

import (
	"database/sql"
	"encoding/json"

	//Can't add these just yet because goland gets rid of them when not being used
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var mismatchedColorsdb *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	mismatchedColorsdb, err = sql.Open("sqlite3", "../database/MismatchedColors.db")
	if err != nil {
		log.Fatal(err)
	}
	defer mismatchedColorsdb.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	router.HandleFunc("/getmismatchedcolors", getMismatchedColorsHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func getMismatchedColorsHandler(w http.ResponseWriter, r *http.Request) {
	colors := colorDB.GetAllRandomColors(mismatchedColorsdb)

	colorJSON, err := json.Marshal(colors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(colorJSON)
}
