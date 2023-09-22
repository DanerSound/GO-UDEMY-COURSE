package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fileWriter struct{}

func (fileWriter) Read (bs []byte)(int, error){
	fmt.Println(string(bs))
	fmt.Println("Just write all this Bytes", len(bs))
	return len(bs),nil
}



func main() {

	fcont, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fw:= fileWriter{}

	io.Copy(fw,fcont.)

}
