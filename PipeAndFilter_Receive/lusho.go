package main

import (
	"fmt"
	"io"
	"net/http"
)
var lusho string="a"

func hello(w http.ResponseWriter, r *http.Request) {
	lusho+="a\n"
	io.WriteString(w, lusho)
	fmt.Println(lusho)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
