package main

import (
	"encoding/json"
	//"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/L0rd5n0w/bookstore/models"
)

const dbFile = "db.json"
var (
	newbook []*models.Books
)

func NewBooks(title, description, author, edition string) *models.Books {
	return &models.Books{
		Title: title,
		Description: description,
		Author: author,
		Edition: edition,
	}
}

func(app *application) home(w http.ResponseWriter, r *http.Request) {
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

func(app *application) formhandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	author := r.FormValue("author")
	edition := r.FormValue("edition")

	newbook = append(newbook, NewBooks(title, description, author, edition))

	// Parsing the form into json
	js, err := json.MarshalIndent(newbook, "", " ")
	if err != nil {
		log.Printf("problem started right here %+v", err)
	}
	
	if len(newbook) > 0 {
		save(js)
		//fmt.Print(string(js))
	}

	w.Write([]byte("Saving to database"))
}

func(app *application) form(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/html/form.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
		//http.Error(w, "Internal Server Error", 500)
	}
}
