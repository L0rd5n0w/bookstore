package main

import (
	"net/http"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}
	//w.Write([]byte("Loading..."))
}

func form(w http.ResponseWriter, _ *http.Request) {
	ts, err := template.ParseFiles("./ui/html/form.html")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Errpr", 500)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Uploading..."))
}