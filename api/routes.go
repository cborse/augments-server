package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mw := alice.New(app.recoverPanic, app.logRequest, app.setHeaders)
	auth := alice.New(app.auth)

	mux := pat.New()

	mux.Post("/login", http.HandlerFunc(app.login))

	mux.Post("/assign", auth.ThenFunc(app.assign))
	mux.Post("/unassign", auth.ThenFunc(app.unassign))
	mux.Post("/hatch_egg", auth.ThenFunc(app.hatchEgg))

	return mw.Then(mux)
}
