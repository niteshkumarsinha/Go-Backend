# Chapter 3: Static File Server

This chapter demonstrates how to serve static assets (HTML, CSS, JS, images) using Go's standard library.

## Key Concepts

### `http.FileServer`

A helper function that returns a handler that serves HTTP requests with the contents of the file system.

### `http.Dir`

A type that implements the `http.FileSystem` interface using the native file system, restricted to a specific directory tree.

### `http.StripPrefix`

When serving files from a subdirectory (e.g., `/static/`), we often need to strip that prefix from the URL path before looking up the file in the file system. `http.StripPrefix` wraps a handler and removes the specified prefix from the request URL.

## Code Explanation

```go
func main(){
    // Create a file server for the "./static" directory
	fs := http.FileServer(http.Dir("./static"))
    
    // Handle requests to "/static/"
    // Strip "/static/" so requesting "/static/style.css" looks for "./static/style.css"
	http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Serve a specific file (index.html) for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "./static/index.html")
	})
    
	http.ListenAndServe(":8080", nil)
}
```
