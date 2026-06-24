// helper file for minor but important
// helper functions
package main

import (
	"log"
	"os"
)

//save: an helper to save the marshalled form input
// and also check for existing books before saving
// into "dbFile"
func save(instantBook []byte) {
	file, err := os.OpenFile(dbFile, os.O_CREATE | os.O_RDWR, 0664)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	rFile, err := os.ReadFile(dbFile)
	if err != nil {
		log.Print(err)
	}
	if len(rFile) > 0 {
		
	} 
	
	file.Write(instantBook)
}