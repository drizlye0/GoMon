package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s in: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s in: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, "bad request")
}
