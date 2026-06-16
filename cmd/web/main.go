package main

import (
	"github.com/L0rd5n0w/bookstore/models"
	"log"
	"net/http"
)

type application struct {
	Books	*models.Books
}

func main() {
	
	app := &application{
		Books: &models.Books{},
	}

	log.Print("Starting Server on :8000")
	err :=http.ListenAndServe(":8000", app.routes())
	log.Fatal(err)

}