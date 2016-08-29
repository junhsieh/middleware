package middleware

import (
	"log"
	"net/http"
)

// Logger ...
type Logger struct {
	handler http.Handler
}

// NewLogger ...
func NewLogger(h http.Handler) *Logger {
	return &Logger{handler: h}
}

// LoggerHandler ...
func LoggerHandler(h http.Handler) http.Handler {
	return NewLogger(h)
}

// Any struct with the method ServeHTTP(http.ResponseWriter, *http.Request) will be
// implementing http.Handler and will be usable with the Go muxer (http.Handle(pattern, handler) function).
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//host := r.Host
	//addr := r.RemoteAddr
	//url := r.URL

	// r.URL.Scheme will be empty if you're accessing the HTTP server not from an HTTP proxy,
	// a browser can issue a relative HTTP request instead of a absolute URL.
	// Additionally, you could check in the server/handler whether you get a
	// relative or absolute URL in the request by calling the IsAbs() method.
	// Reference: http://stackoverflow.com/questions/6899069/why-are-request-url-host-and-scheme-blank-in-the-development-server
	//scheme := r.URL.Scheme
	//isAbs := r.URL.IsAbs()

	//uri := r.RequestURI

	//log.Printf("Before: %v; %v; %v; %v", host, addr, url, uri)
	log.Printf("DEBUG_LOGGER: Inside")

	l.handler.ServeHTTP(w, r)

	//log.Printf("After: %v\t%v", r.Host, r.URL.Path)
}
