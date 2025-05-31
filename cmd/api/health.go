package main

import "net/http"

func (app *application) checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.jsonResponse(w, http.StatusOK, "status OK"); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
