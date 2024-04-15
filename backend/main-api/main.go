package main

import (
	//Can't add these just yet because goland gets rid of them when not being used
	//"mismatchedColors/models"
	//"mismatchedColors/databse"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprint(w, "Hello, World!")
	// })
	router.HandleFunc("/", hello).Methods(http.MethodPost)
	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
