package middleware

import (
	"github.com/junhsieh/iojson"
	"net/http"
)

func DomainHandler(allowedDomain string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Host == allowedDomain {
				h.ServeHTTP(w, r)
			} else {
				o := r.Context().Value("iojson").(*iojson.IOJSON)
				o.AddError("Invalid domain")
			}
		})
	}
}
