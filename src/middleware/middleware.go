package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// Adapter asd
type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LogRequest logs information about the incoming request.
func LogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)

		defer func() {
			start := time.Now()
			duration := time.Since(start)
			logrus.WithFields(logrus.Fields{
				"status":           lrw.statusCode,
				"method":           r.Method,
				"route":            r.URL.String(),
				"request-duration": duration,
				"remote":           r.RemoteAddr,
				"referer":          r.Referer(),
				"user-agent":       r.UserAgent(),
			}).Info("completed request")
		}()

		h.ServeHTTP(lrw, r)
	})
}

// PanicRecovery gracefully handles panics in application code by returning a 500 error.
func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
