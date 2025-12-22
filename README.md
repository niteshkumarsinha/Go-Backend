# Go Backend Learning Path

Welcome to the Go Backend Learning repository! This project serves as a step-by-step guide to mastering backend development with Go (Golang). Each chapter introduces core concepts, building upon the previous ones to help you construct robust and efficient web servers.

## Table of Contents

- [Chapter 1: Introduction to HTTP Server](#chapter-1-introduction-to-http-server)
- [Chapter 2: Handlers & ServeMux](#chapter-2-handlers--servemux)
- [Chapter 3: Static File Server](#chapter-3-static-file-server)
- [Chapter 4: Query Parameters & Path variables](#chapter-4-query-parameters--path-variables)
- [Chapter 5: Middlewares](#chapter-5-middlewares)
- [Chapter 6: Building a JSON API](#chapter-6-building-a-json-api)

---

## Chapter 1: Introduction to HTTP Server

**Directory:** `src/1-intro`

This chapter covers the absolute basics of setting up an HTTP server in Go using the standard `net/http` library.

### Key Concepts
- **`http.HandleFunc`**: Registers a function to handle requests to a specific pattern (route).
- **`http.ListenAndServe`**: Starts the HTTP server on a specified address and port.
- **`http.ResponseWriter`**: Used to construct and send the HTTP response.
- **`*http.Request`**: detailed information about the incoming request.

### Example
```go
func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func main() {
    http.HandleFunc("/api", apiHandler)
    http.ListenAndServe(":8080", nil) // nil uses the DefaultServeMux
}
```

---

## Chapter 2: Handlers & ServeMux

**Directory:** `src/2-handlers`

Here we dive deeper into how Go handles HTTP requests using the `Handler` interface and request multiplexing.

### Key Concepts
- **`http.Handler` Interface**: Any type that implements `ServeHTTP(ResponseWriter, *Request)` can be a handler.
- **`http.NewServeMux()`**: Creating a custom request multiplexer (router) instead of using the global default one.
- **Custom Struct Handlers**: Attaching `ServeHTTP` to a struct allows holding dependencies or state.

### Example
```go
type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Go Server")
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", HomeHandler{}) // Registering a struct handler
    http.ListenAndServe(":8080", mux)
}
```

---

## Chapter 3: Static File Server

**Directory:** `src/3-static-file-server`

Learn how to serve static assets like HTML, CSS, and images.

### Key Concepts
- **`http.FileServer`**: A built-in handler that serves files from a specific directory.
- **`http.Dir`**: Specifies the file system directory to serve.
- **`http.StripPrefix`**: Necessary when serving files under a specific route (e.g., `/static/`) to map the URL path correctly to the file system path.
- **`http.ServeFile`**: Helper to serve a single specific file.

### Example
```go
func main() {
    fs := http.FileServer(http.Dir("./static"))
    // Maps /static/style.css -> ./static/style.css
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Serve a specific index file for root
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "./static/index.html")
    })
}
```

---

## Chapter 4: Query Parameters & Path variables

**Directory:** `src/4-query-path`

Interacting with dynamic data from the URL is crucial for APIs.

### Key Concepts
- **Query Parameters**: Accessed via `r.URL.Query()`. Example: `?name=John`.
- **Path Parameters**: Go's standard library (pre-1.22) doesn't support named path parameters (like `/user/:id`) out of the box. We handle this by manually parsing `r.URL.Path`.

### Example
```go
// Query Params: /greet?name=John
func greetHandler(w http.ResponseWriter, r *http.Request){
    name := r.URL.Query().Get("name")
    if name == "" { name = "Guest" }
    fmt.Fprintln(w, "Hello,", name)
}

// Path Params: /api/v1/user/123
func userHandler(w http.ResponseWriter, r *http.Request){
    // Manual parsing
    segments := strings.Split(r.URL.Path, "/")
    userId := segments[len(segments)-1]
    fmt.Fprintf(w, "User ID: %s", userId)
}
```

---

## Chapter 5: Middlewares

**Directory:** `src/5-middlewares`

Middleware allows you to intercept and process requests before they reach your final handler. Useful for logging, authentication, and setting headers.

### Key Concepts
- **Middleware Pattern**: A function that takes an `http.Handler` and returns a new `http.Handler`.
- **Chaining**: Wrapping handlers in multiple middlewares (e.g., `Logging(Auth(Handler))`).

### Example
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request received:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    finalHandler := http.HandlerFunc(homeHandler)
    http.Handle("/", loggingMiddleware(finalHandler))
}
```

---

## Chapter 6: Building a JSON API

**Directory:** `src/6-json-API`

Pulling it all together to build a RESTful JSON API.

### Key Concepts
- **`encoding/json`**: Used for parsing incoming JSON (`Decode`) and sending JSON responses (`Encode`).
- **HTTP Methods**: Switching logic based on `r.Method` (GET, POST, PUT, DELETE).
- **Thread Safety**: Using `sync.Mutex` to safely manage concurrent access to shared data (like an in-memory map).

### Example
```go
type User struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(currentUser)
    case http.MethodPost:
        var newUser User
        json.NewDecoder(r.Body).Decode(&newUser)
        // ... save user
        w.WriteHeader(http.StatusCreated)
    }
}
```
