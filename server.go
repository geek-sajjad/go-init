package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, request *http.Request) {
		io.WriteString(w, "Hello world")
	}

	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting the server:", err)

	}
}
