package main

import (
	"log"
	"os"
)

func main() {
	path := "test.txt"
	newPath := "testRenamed.txt"

	err := os.Rename(path, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
