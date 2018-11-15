package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server
	fmt.Println("Running web services on port 8080.....")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
