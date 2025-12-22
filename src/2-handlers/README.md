# Chapter 2: Handlers & ServeMux

In this chapter, we explore how to gain more control over request routing and handling by using custom Handlers and a ServeMux.

## Key Concepts

### `http.Handler` Interface

In Go, any type that implements the `ServeHTTP` method is an `http.Handler`. This allows you to create stateful handlers (e.g., passing database connections).

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### `ServeMux`

A `ServeMux` is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

- **`http.NewServeMux()`**: Creates a new, isolated multiplexer instance. This is safer than using the global default mux in larger applications.

## Code Explanation

```go
// Custom Handler Struct
type HomeHandler struct{}

// Implement ServeHTTP to satisfy the Handler interface
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go Server")
}

func main() {
	mux := http.NewServeMux()
    
    // Register struct handler
	mux.Handle("/", HomeHandler{})
    
    // Register functional handler
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
    
    // Use the custom mux
	http.ListenAndServe(":8080", mux)
}
```
