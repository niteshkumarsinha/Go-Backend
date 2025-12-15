package main

import (
	"fmt"
	"net/http"
)

// *http.Request -> points to location where user requests and parameters are present i.e User Provided Data
// http.ResponseWriter -> Backend writes its response
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// "Hello World" -> w
	fmt.Fprintln(w, "Hello World")
}

func main() {
	// localhost:8080/api -> invokes handler function when this path is called
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/user", apiHandler)

	fmt.Println("Starting server on port 8080...")
	// host:port -> localhost:8080 | :8080 -> listen and serve at all interfaces on port 8080
	http.ListenAndServe(":8080", nil)
}
