package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("sah dude")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, %s", html.EscapeString(request.URL.Path))
		fmt.Fprintf(writer, "This came from the web server")
	})

	log.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
