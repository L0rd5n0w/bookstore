package main

import (
	"net/http"
	"log"
)

func main() {
	
	log.Print("Starting Server on :8000")
	err :=http.ListenAndServe(":8000", routes())
	log.Fatal(err)

}