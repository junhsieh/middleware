package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/junhsieh/iojson"
	"log"
	"net/http"
)

// AuthUserHandler ...
func AuthUserHandler(store *sessions.FilesystemStore, sessionName string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("DEBUG_AuthUser: Inside")

			o := r.Context().Value("iojson").(*iojson.IOJSON)

			// Get a session. Get() always returns a session, even if empty.
			session, err := store.Get(r, sessionName)

			if err != nil {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				//w.WriteHeader(http.StatusInternalServerError)
				o.AddError(err.Error())
				return
			}

			// TODO: how to determine whether a session is expired?

			// TODO: add a logic to continue only when session.IsNew is false (can be true too when the session is expired)
			log.Printf("DEBUG_AuthUser: IsNew Session: %v", session.IsNew)

			// TODO: need a way to check if session exists.

			if _, ok := session.Values["Username"]; !ok {
				//w.WriteHeader(http.StatusForbidden)
				o.AddError("You do not have the permission")
				return
			}

			o.AddData("welcome", "Welcom, "+session.Values["Username"].(string))

			h.ServeHTTP(w, r)

			// TODO: Add a logic to support both session/cookie and Header/Authorization
			/*
				Authorization := r.Header.Get("Authorization")

				log.Printf("Authorization: %v", Authorization)

				switch Authorization {
				case "user2":
					w.WriteHeader(http.StatusForbidden)
					o.AddError("You do not have the permission")
				case "user3":
					o.AddError("You do not have the permission")
				}
			*/
		})
	}
}
