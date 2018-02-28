package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	app "github.com/lwalter/lessonshare-api/src/app"
	"github.com/lwalter/lessonshare-api/src/data"
	"github.com/lwalter/lessonshare-api/src/handlers"
	"github.com/lwalter/lessonshare-api/src/middleware"
)

func initAPI() *mux.Router {
	router := mux.NewRouter()

	ns := "api"
	v1 := "v1"

	// Routes
	pingHandler := http.HandlerFunc(handlers.Ping)
	getLessonsHandler := http.HandlerFunc(handlers.GetLessons)
	router.Handle("/ping", middleware.LogRequest(pingHandler)).Methods("GET")
	router.Handle("/"+ns+"/"+v1+"/lessons", middleware.LogRequest(getLessonsHandler)).Methods("GET")
	router.HandleFunc("/"+ns+"/"+v1+"/lessons/{id}", handlers.GetLesson).Methods("GET")

	return router
}

func main() {
	app.LoadConfig()
	data.InitDbConn()
	app.InitLog()
	router := initAPI()
	p := strconv.Itoa(app.Config.App.Port)

	if err := http.ListenAndServe(":"+p, router); err != nil {
		log.Fatal(err)
	}
}
