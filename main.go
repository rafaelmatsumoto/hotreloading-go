package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Listening on port 8000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Pather = %q\n", r.URL.Path)
}
