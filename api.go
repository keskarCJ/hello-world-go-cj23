package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	t "time"
)

type Time struct {
	Now string `json:"now,omitempty"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	time := Time{
		Now: t.Now().String(),
	}

	fmt.Printf("Get request from %v", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(time)
}

func CreateTime(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	log.Printf("%v", params)

	var time Time

	json.NewDecoder(r.Body).Decode(&time)

	log.Printf("%v\n", time)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/post", CreateTime).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}