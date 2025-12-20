package main


import (
	"fmt"
	"log"
	"net/http"
	//"time"
)

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "Go Server")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received for", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to the about page!")
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))
	mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

	log.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		//panic(err)
		log.Fatal("Server Failed to start", err)
	}
}


