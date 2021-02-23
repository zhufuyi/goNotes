package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "say hello", r.Host, r.URL.Path)
}

func main() {
	http.HandleFunc("/sayHello", sayHello)
	http.ListenAndServe(":8080", nil)
}
