package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
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
	pingHandler := middleware.Adapt(http.HandlerFunc(handlers.Ping), middleware.PanicRecovery, middleware.LogRequest)
	getLessonsHandler := middleware.Adapt(http.HandlerFunc(handlers.GetLessons), middleware.PanicRecovery, middleware.LogRequest)
	getLessonHandler := middleware.Adapt(http.HandlerFunc(handlers.GetLesson), middleware.PanicRecovery, middleware.LogRequest)
	router.Handle("/ping", pingHandler).Methods("GET")
	router.Handle("/"+ns+"/"+v1+"/lessons", getLessonsHandler).Methods("GET")
	router.Handle("/"+ns+"/"+v1+"/lessons/{id}", getLessonHandler).Methods("GET")

	return router
}

func main() {
	app.LoadConfig()
	data.InitDbConn()
	app.InitLog()
	router := initAPI()
	p := strconv.Itoa(app.Config.App.Port)

	address := fmt.Sprintf(":%v", p)
	logrus.Info(fmt.Sprintf("server started at %v\n", address))
	if err := http.ListenAndServe(":"+p, router); err != nil {
		log.Fatal(err)
	}
}
