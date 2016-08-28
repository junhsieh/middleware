package middleware

import (
	"net/http"
)

// Any struct with the method ServeHTTP(http.ResponseWriter, *http.Request) will be
// implementing http.Handler and will be usable with the Go muxer (http.Handle(pattern, handler) function).
type Logger struct {
	handler http.Handler
}

func NewLogger(h http.Handler) *Logger {
	return &Logger{handler: h}
}

func LoggerHandler(h http.Handler) http.Handler {
	return NewLogger(h)
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	addr := r.RemoteAddr
	url := r.URL
	uri := r.RequestURI

	log.Printf("Before: %v; %v; %v; %v", host, addr, url, uri)

	l.handler.ServeHTTP(w, r)

	//log.Printf("After: %v\t%v", r.Host, r.URL.Path)
}
