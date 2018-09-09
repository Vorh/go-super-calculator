package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const urlService = "http://localhost:8080/calculate"

func main() {

	if len(os.Args) > 1 {
		value := os.Args[1]
		sendRequest(value)
	} else {
		fmt.Println("Input expression of the form: x+x")
	}
}

func sendRequest(ex string) {

	resp, err := http.PostForm(urlService, url.Values{"ex": {ex}})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(resp)
}
