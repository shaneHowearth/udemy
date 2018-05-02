package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1:]

	if len(filename) == 0 {
		log.Fatal("No filename specified")
	}
	f, err := os.OpenFile(filename[0], os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}
