package main


import (
	"fmt"
	"net/http"
	"strings"
)

// QUERY PARAMS
// https://api.example.com/api/v1/greet?name=John
func greetHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintln(w, "Hello, ", name)
}

// PATH PARAMS
//https://api.example.com/api/v1/user/1234
// 1 -> User
// 2 -> UserID
func userHandler(w http.ResponseWriter, r *http.Request){
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		fmt.Fprintln(w, "Invalid path")
		return
	}
	if len(pathSegments) >= 3 && pathSegments[len(pathSegments)-2] == "user" {
		userId := pathSegments[len(pathSegments)-1]
		fmt.Fprintf(w, "User ID: %s\n", userId)
	} else {
		http.NotFound(w, r)
	}
}


// COMBINING BOTH QUERY AND PATH PARAMS
// https://api.example.com/username/1234?include_details=true
func userDetailsHandler(w http.ResponseWriter, r *http.Request){
	pathSegments := strings.Split(r.URL.Path, "/")
	query := r.URL.Query()
	includeDetails := query.Get("include_details")
	if len(pathSegments) < 3 {
		fmt.Fprintln(w, "Invalid path")
		return
	}
	if len(pathSegments) >= 3 && pathSegments[len(pathSegments)-2] == "username" {
		userId := pathSegments[len(pathSegments)-1]
		fmt.Fprintf(w, "User ID: %s\n", userId)
		if includeDetails == "true" {
			fmt.Fprintln(w, "Include details: true")
		}
	} else {
		http.NotFound(w, r)
	} 
	
}


func main(){
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/api/v1/user/", userHandler)
	http.HandleFunc("/username/", userDetailsHandler)
	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start at port 8080", err)
		return
	}
}

// https://api.example.com/api/v1/greet?name=John&referral_code=1234
// https://api.example.com/api/v1/greet?name=John&referral_code=1234&session_id=1234
