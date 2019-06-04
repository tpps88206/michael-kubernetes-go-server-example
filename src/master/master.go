package main

import (
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", GetData).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func GetData(w http.ResponseWriter, r *http.Request) {
    resp, _ := http.Get("http://user.default.svc.cluster.local:8080")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    response := []byte(string(body))
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
