package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lwalter/lessonshare-api/src/handlers"
)

func main() {
	router := mux.NewRouter()
	port := 8080

	apiPrefix := "api"
	apiVersion := "v1"

	// Routes
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.HandleFunc("/"+apiPrefix+"/"+apiVersion+"/lessons", handlers.GetLessons).Methods("GET")
	router.HandleFunc("/"+apiPrefix+"/"+apiVersion+"/lessons/{id}", handlers.GetLesson).Methods("GET")

	p := strconv.Itoa(port)
	if err := http.ListenAndServe(":"+p, router); err != nil {
		log.Fatal(err)
	}
}
