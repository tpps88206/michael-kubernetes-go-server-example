package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":39000", router))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var res = map[string]string{"name": "Michael", "email": "michael@gmail.com"}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
