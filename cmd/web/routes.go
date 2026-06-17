package main

import (
	"net/http"
)

func(app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/form", app.form)
	mux.HandleFunc("/form/handler", app.formhandler)

	return mux
}