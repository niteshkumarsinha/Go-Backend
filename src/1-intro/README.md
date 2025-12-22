# Chapter 1: Introduction to HTTP Server

This chapter introduces the fundamental concepts of building an HTTP server in Go using the standard `net/http` package.

## Key Concepts

### `http.HandleFunc`

The `http.HandleFunc` function is used to register a handler function for a specific URL pattern. When a request matches the pattern, the registered function is executed.

### `http.ListenAndServe`

This function starts an HTTP server on a specified network address (specifically TCP) and blocks the process to listen for incoming requests.

### `http.ResponseWriter`

An interface used by the HTTP handler to construct an HTTP response. You use it to write the response body and status codes.

### `*http.Request`

A pointer to a struct that represents the HTTP request received by the server. It contains all the details like URL, headers, and body.

## Code Explanation

```go
// Handler function
func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func main() {
    // Register the handler for the "/api" path
    http.HandleFunc("/api", apiHandler)
    
    // Start the server on port 8080
    http.ListenAndServe(":8080", nil)
}
```

- **`http.HandleFunc("/api", apiHandler)`**: Tells the server to call `apiHandler` whenever a request is made to `/api`.
- **`nil` in `ListenAndServe`**: Indicates that we are using the `DefaultServeMux` (the default router).
