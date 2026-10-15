package main

import (
	"fmt"
	"net/http"

	"example.com/go-lez14/hi"
)

func main() {
	http.HandleFunc("/", hi.SayHi2)
	fmt.Println("Starting web server at http://localhost:8090")
	http.ListenAndServe(":8090", nil)

}
