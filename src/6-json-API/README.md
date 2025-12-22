# Chapter 6: Building a JSON API

This final chapter combines previous concepts to build a fully functional RESTful JSON API with CRUD operations.

## Key Concepts

### RESTful Methods

We switch behavior based on the HTTP method (`GET`, `POST`, `PUT`, `DELETE`) to perform different actions on a resource.
- `r.Method` gives us the method string (e.g., "GET").

### JSON Encoding/Decoding

- **`json.NewEncoder(w).Encode(data)`**: Converts Go structs into JSON and writes them to the response.
- **`json.NewDecoder(r.Body).Decode(&data)`**: Reads the request body stream and unmarshals it into a Go struct.

### Concurrency Safety (`sync.Mutex`)

Since HTTP servers in Go handle requests concurrently (in separate goroutines), accessing shared memory (like a global map) requires synchronization to prevent race conditions.
- **`mutex.Lock()` / `mutex.Unlock()`**: Ensures only one goroutine accesses the map at a time.

## Code Explanation

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
        // Read (GET)
		json.NewEncoder(w).Encode(user)
        
	case http.MethodPut:
        // Update (PUT)
		var updatedUser User
		json.NewDecoder(r.Body).Decode(&updatedUser)
        // ... update logic
	}
}
```
