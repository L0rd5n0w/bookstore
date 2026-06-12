package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/form", form)

	return mux
}