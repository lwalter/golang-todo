package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lwalter/lessonshare-api/src/data"
)

func GetLessons(w http.ResponseWriter, r *http.Request) {
	lessons, err := data.GetAllLessons()
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	if lessons == nil || len(lessons) == 0 {
		http.Error(w, "Not found", 404)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(lessons)
}

func GetLesson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "ID value is invalid", 400)
		return
	}

	lesson, err := data.GetLesson(id)

	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	if lesson != nil {
		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode(lesson)
	} else {
		http.Error(w, "Not found", 404)
	}
}
