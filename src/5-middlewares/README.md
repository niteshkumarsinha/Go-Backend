# Chapter 5: Middlewares

Middleware is a powerful pattern in backend development for separating cross-cutting concerns (like logging, authentication, or error handling) from your business logic.

## Key Concepts

### Middleware Pattern

A middleware is simply a function that takes an `http.Handler` and returns a new `http.Handler` that wraps the original one. It can perform actions *before* and *after* the wrapped handler is executed.

```go
func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Pre-processing
        next.ServeHTTP(w, r)
        // Post-processing
    })
}
```

### Chaining

You can chain multiple middlewares together to create a processing pipeline.
`MiddlewareA(MiddlewareB(FinalHandler))`

## Code Explanation

```go
// Logging Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received for", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Header Middleware
func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
    // Chain: logging -> header -> homeHandler
	mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))
}
```
