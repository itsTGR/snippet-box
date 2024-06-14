package main

import (
	"net/http"
	"runtime/debug"
)

// Method serverError writes an error entry at Error level and sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Method clientError sends a specific status code and its corresponding description to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Method NotFound is a wrapper for clientError, where status code to use will be 404
func (app *application) NotFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
