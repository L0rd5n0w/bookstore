package main

import (
	"net/http"
)

func(app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/form", app.form)
	mux.HandleFunc("/upload", upload)

	return mux
}