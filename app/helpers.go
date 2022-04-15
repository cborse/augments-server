package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s/n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter) {
	trace := fmt.Sprintf("%s/n%s", http.StatusText(http.StatusBadRequest), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (app *application) writeStruct(w http.ResponseWriter, v interface{}) {
	resp, err := json.Marshal(v)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func getCredentials(r *http.Request) (string, uint64) {
	// Format
	// X-Aug-ID: xxx
	// X-Aug-Token: xxx
	userID, err := strconv.ParseUint(r.Header.Get("X-Aug-ID"), 10, 32)
	if err != nil || userID == 0 {
		return "", 0
	}
	token := r.Header.Get("X-Aug-Token")
	return token, userID
}
