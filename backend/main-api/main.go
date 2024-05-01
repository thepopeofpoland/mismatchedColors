package main

import (
	"mismatchedColors/database"
	"mismatchedColors/models"

	"database/sql"
	"encoding/json"

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
	router.HandleFunc("/getrandomcolors", getRandomColorsHandler).Methods(http.MethodPost)
	router.HandleFunc("/insertrandomcolor", insertRandomColorHandler).Methods(http.MethodPost)

	router.HandleFunc("/getcolorpalettes", getColorPalettes).Methods(http.MethodPost)
	router.HandleFunc("/insertcolorpalette", insertColorPaletteHandler).Methods(https.MethodPost)

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func getRandomColorsHandler(w http.ResponseWriter, r *http.Request) {
	colors := colorDB.GetAllRandomColors(mismatchedColorsdb)

	colorJSON, err := json.Marshal(colors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(colorJSON)
}

func insertRandomColorHandler(w http.ResponseWriter, r *http.Request) {
	var randomColor models.RandomColors

	err := json.NewDecoder(r.Body).Decode(&randomColor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Put some error handling here to check for missing parameters (bad JSON)

	newColor := colorDB.InsertNewRandomColor(mismatchedColorsdb, randomColor)

	newSupplyJson, err := json.Marshal(newColor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(newSupplyJson)
}


// --- Color Palettes ---

func getColorPalettes(w http.ResponseWriter, r *http.Request) {
	palettes := colorDB.GetAllColorPalettes(mismatchedColorsdb)

	paletteJSON, err := json.Marshal(palettes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(paletteJSON)
}


func insertColorPaletteHandler(w http.ResponseWriter, r *http.Request) {
	var colorPalette models.ColorPalettes

	err := json.NewDecoder(r.Body).Decode(&colorPalette)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Put some error handling here to check for missing parameters (bad JSON)

	newPalette := colorDB.InsertNewColorPalette(mismatchedColorsdb, colorPalette)

	newSupplyJson, err := json.Marshal(newPalette)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(newSupplyJson)
}