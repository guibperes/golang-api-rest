package log

import (
	"log"
	"net/http"
)

// Log definition and middleware creation
type Log struct{}

// Builder build the Log object
func Builder() Log {
	return Log{}
}

// Info log type
func (l Log) Info(value interface{}) {
	log.Println("[INFO] " + value.(string))
}

// Fatal log type
func (l Log) Fatal(value interface{}) {
	log.Fatalln("[FATAL] " + value.(string))
}

// GetLogMiddleware return logging middleware
func (l Log) GetLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Info(r.Method + " " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
