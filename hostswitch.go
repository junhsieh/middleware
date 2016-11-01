package middleware

import (
	"net/http"
)

type HostSwitch map[string]http.Handler

func New() HostSwitch {
	return make(HostSwitch)
}

// Implement the ServerHTTP method
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler, ok := hs[r.Host]; handler == nil || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		handler.ServeHTTP(w, r)
	}
}
