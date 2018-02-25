package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	cfg "github.com/lwalter/lessonshare-api/src/config"
	"github.com/lwalter/lessonshare-api/src/data"
	"github.com/lwalter/lessonshare-api/src/handlers"
)

func initApi() *mux.Router {
	router := mux.NewRouter()

	apiNs := "api"
	apiV := "v1"

	// Routes
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.HandleFunc("/"+apiNs+"/"+apiV+"/lessons", handlers.GetLessons).Methods("GET")
	router.HandleFunc("/"+apiNs+"/"+apiV+"/lessons/{id}", handlers.GetLesson).Methods("GET")

	return router
}

func main() {
	cfg.LoadConfig()
	data.InitDbConn()
	router := initApi()
	p := strconv.Itoa(cfg.Config.App.Port)

	if err := http.ListenAndServe(":"+p, router); err != nil {
		log.Fatal(err)
	}
}
