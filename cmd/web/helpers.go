package main

import (
	"os"
	"log"
)

func save(instantBook []byte) {
	file, err := os.OpenFile(dbFile, os.O_CREATE | os.O_RDWR, 0664)
	if err != nil {
		log.Print(err)
	}
	file.Write(instantBook)

	file.Close()
}