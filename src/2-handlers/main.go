package main

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go Server")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", HomeHandler{})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! You are accessing %s and using user agent %s\n", r.URL.Path, r.Header.Get("User-Agent"))
	})
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "ok"}`)
	})
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", mux)
}
