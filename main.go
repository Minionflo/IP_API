package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, r.Header.Get("x-forwarded-for"))
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
