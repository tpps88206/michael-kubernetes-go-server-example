package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var userServiceUrl string

func main() {
	userServiceUrl = os.Getenv("USER_SERVICE_PATH")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetData).Methods("GET")

	log.Fatal(http.ListenAndServe(":39000", router))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(userServiceUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := []byte(string(body))
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
