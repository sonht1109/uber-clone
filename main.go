package main

import (
	"fmt"
	"net/http"
)

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/ping", getData)
	http.ListenAndServe(":81", nil)
}
