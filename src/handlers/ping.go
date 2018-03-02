package handlers

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	panic("testing panic handler")
	w.WriteHeader(http.StatusOK)
}
