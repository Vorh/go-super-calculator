package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/calculate", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	get := r.PostForm.Get("ex")
	fmt.Println(get)
}
