package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello HTTP SERVER")
}

func Text(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello How are you Today")
}
func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/text", Text)
	fmt.Println("Server : https://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
