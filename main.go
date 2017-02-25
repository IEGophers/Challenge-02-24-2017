package main

import (
	"flag"
	"io"
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

	tr, err := createTarball(*path)
	if err != nil {
		log.Fatal(err)
	}

}

func createTarball(path string) (io.Reader, error) {
	// your code here
	return nil, nil
}
