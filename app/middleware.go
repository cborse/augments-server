package main

import (
	"augments/models"
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

func (app *application) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the credentials
		token, userID := getCredentials(r)
		if len(token) < 1 || userID == 0 {
			app.clientError(w, http.StatusUnauthorized)
			return
		}

		// Find the user
		user := &models.User{}
		err := app.db.Get(user, "SELECT * FROM user WHERE id = ?", userID)
		if err != nil {
			app.serverError(w, err)
			return
		}
		if user.Token != token {
			app.clientError(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
