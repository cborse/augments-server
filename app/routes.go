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
	mux.Get("/get_data", auth.ThenFunc(app.getData))
	mux.Post("/assign", auth.ThenFunc(app.assign))
	mux.Post("/unassign", auth.ThenFunc(app.unassign))
	mux.Post("/hatch_egg", auth.ThenFunc(app.hatchEgg))
	mux.Post("/learn_action", auth.ThenFunc(app.learnAction))
	mux.Post("/learn_skill", auth.ThenFunc(app.learnSkill))
	mux.Post("/matchmake", auth.ThenFunc(app.matchmake))

	return mw.Then(mux)
}
