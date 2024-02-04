package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/hello-world", hello)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<div hx-get=\"/hello-world\">Click me!</div>")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<div>Hello, world!</div>")
}
