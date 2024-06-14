package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.NotFound(w)
		return // Don't forget the 'return'
	}

	files := []string{
		"./ui/html/base.tmpl.htmll", //Base template needs to be the first in this list
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html", // If some template file needs to be added, just add it to this slice
	}

	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		//app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		app.serverError(w, r, err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		//errMsg := fmt.Sprintf("Invalid ID: %d\nERR: %s\n", id, err.Error())
		//w.Write([]byte(errMsg))
		//http.NotFound(w, r)
		app.NotFound(w)
		return
	}

	//msg := fmt.Sprintf()
	//w.Write([]byte(msg))

	fmt.Fprintf(w, "Displaying snippet %d\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//w.Header()["from-tgr"] = []string{"1; mode=block"}
		//w.Header()["X-Content-Type-Options"] = nil
		//w.Header()["Date"] = nil
		//w.Header().Add("Allow", http.MethodGet)
		//fmt.Println(w.Header().Get("Allow"))
		//fmt.Println(w.Header().Values("Allow"))
		//w.Header().Del("Allow")
		//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) //http.Error is a helper function that calls w.Write and w.WriteHeader under the hood
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed\n"))
		w.Header().Set("Allow", http.MethodPost) // w.Header().Set will not take effect in case it is after w.WriteHeader or w.Write
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a new snippet\n"))
}
