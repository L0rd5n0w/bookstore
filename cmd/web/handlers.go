package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
//	"os"
)

const dbFile = "db.json"

type Books struct {
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Author		string	`json:"author"`
	Edition		string	`json:"edition"`
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

func(app *application) form(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/html/form.gohtml")
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	author := r.FormValue("author")
	edition := r.FormValue("edition")


	newbook := Books{
		Title: title,
		Description: description,
		Author: author,
		Edition: edition,
	}

	

	// for parsing the form into json and
	// storing it into a json file
	js, err := json.MarshalIndent(newbook, "", " ")
	if err != nil {
		log.Printf("problem started right here %+v", err)
	}

	err = ts.Execute(w, newbook)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
	}

	fmt.Print(string(js))
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Uploading..."))
}