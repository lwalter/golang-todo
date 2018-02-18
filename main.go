package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lwalter/lessonshare-api/src/handlers"
)

func main() {
	router := mux.NewRouter()
	port := 8080

	apiPrefix := "api"
	apiVersion := "v1"

	// Routes
	router.HandleFunc(apiPrefix+apiVersion+"/lessons", handlers.GetLessons).Methods("GET")
	router.HandleFunc(apiPrefix+apiVersion+"/lessons/{id}", handlers.GetLesson).Methods("GET")

	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		log.Fatal(err)
	}
}
