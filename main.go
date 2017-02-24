package main

import (
	"flag"
	"log"
)

var (
	path = flag.String("path", "", "path to tar")
	out  = flag.String("out", "", "file to create")
)

func main() {
	if *path == "" {
		log.Fatal("path is a required argument.")
	}

	if *out == "" {
		log.Fatal("out is a required argument.")
	}

	if err := createTarball(*path, *out); err != nil {
		log.Fatal(err)
	}
}

func createTarball(path, out string) error {
	// your code here
	return nil
}
