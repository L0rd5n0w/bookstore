package main

import (
	"html/template"
	"log"
	"net/http"
)

type Books struct {
	Title		string	`form:"title"`
	Description	string	`form:"description"`
	Author		string	`form:"author"`
	Edition		string	`form:"edition"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/html/home.gohtml")
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
}

func form(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/html/form.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	author := r.FormValue("author")
	edition := r.FormValue("edition")

	// an instantiation of Books
	iBook := Books{
		Title: title,
		Description: description,
		Author: author,
		Edition: edition,
	}

	err = ts.Execute(w, iBook)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Uploading..."))
}