// Server4 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
  "go_book/ch1/lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("Server listening on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
 lissajous.Lissajous(w) 
}
