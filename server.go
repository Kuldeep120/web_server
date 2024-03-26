package main

import (
	"fmt"
	"log"
	"net/http"
)

// import "log"
type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	// http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Print("random test string")
}
