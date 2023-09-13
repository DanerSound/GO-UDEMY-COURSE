package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write (bs []byte)(int, error){
	fmt.Println(string(bs))
	fmt.Println("Just write all this Bytes", len(bs))
	return len(bs),nil
}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	/* This is in general , but it can be semplify , se next session
	bs := make([]byte, 99999) // empty holder that will contain the content that will arrive
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	*/
	io.Copy(os.Stdout, resp.Body)

}
  