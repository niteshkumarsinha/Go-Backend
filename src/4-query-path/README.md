# Chapter 4: Query Parameters & Path Variables

This chapter focuses on extracting dynamic data from the URL, both through query strings and path segments.

## Key Concepts

### Query Parameters

Query parameters appear after the `?` in a URL (e.g., `/greet?name=Alice`).
- **`r.URL.Query()`**: Returns a map of query values.
- **`.Get("key")`**: Retrieves the first value associated with the given key.

### Path Parameters

Path parameters are part of the URL path itself (e.g., `/user/123`).
- **Standard Library (pre-1.22)**: Standard `net/http` didn't fully support named path parameters like `/user/{id}` until Go 1.22.
- **Manual Parsing**: We can manually inspect `r.URL.Path` by splitting it with `strings.Split` to extract segments.

## Code Explanation

```go
// Handling Query Params
func greetHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	name := query.Get("name") // Get ?name=...
    // ...
}

// Handling Path Params manually
// URL: /api/v1/user/1234
func userHandler(w http.ResponseWriter, r *http.Request){
	pathSegments := strings.Split(r.URL.Path, "/")
    // pathSegments -> ["", "api", "v1", "user", "1234"]
	
    // Extract ID from the last segment
    userId := pathSegments[len(pathSegments)-1]
	fmt.Fprintf(w, "User ID: %s", userId)
}
```
