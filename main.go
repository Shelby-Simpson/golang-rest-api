package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World thanks for connecting to me.")
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler 1")
}

func Now(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func NewYork(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("US/Eastern")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func London(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/us", handler1).Methods("GET")
	r.HandleFunc("/now", Now).Methods("GET")
	r.HandleFunc("/newyork", NewYork).Methods("GET")
	r.HandleFunc("/london", London).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
