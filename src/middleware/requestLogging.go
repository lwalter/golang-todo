package middleware

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// LogRequest logs information about the incoming request.
func LogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"method":           r.Method,
			"route":            r.URL.String(),
			"request-duration": duration,
			"remote":           r.RemoteAddr,
			"referer":          r.Referer(),
			"user-agent":       r.UserAgent(),
		}).Info("completed request")
	})
}
