// helper file for minor but important
// helper functions
package main

import (
	"os"
	"log"
)

//save: an helper to save the marshalled form input 
// into a "dbFile"
func save(instantBook []byte) {
	file, err := os.OpenFile(dbFile, os.O_CREATE | os.O_RDWR, 0664)
	if err != nil {
		log.Print(err)
	}
	file.Write(instantBook)

	file.Close()
}