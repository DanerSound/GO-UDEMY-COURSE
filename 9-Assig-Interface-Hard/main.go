package main

import (
	"io"
	"log"
	"os"
)

func main() {

	fcont, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fcont)

}
