package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email at <a href=\"mailto:devin@duclayan.com\">"+
		"devin@duclayan.com</a>.")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", router)
}
