// helper file for minor but important
// helper functions
package main

import (
	"log"
	"os"
)

//save: an helper to save the marshalled form input
// into a "dbFile"
func save(instantBook []byte) {
	file, err := os.OpenFile(dbFile, os.O_CREATE | os.O_RDWR, 0664)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	file.Write(instantBook)
}