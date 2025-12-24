# Chapter 7: Structs, Methods, and Data Validation

**Directory:** `src/7-structs-validations-error-handling`

This chapter focuses on defining data structures, attaching methods for behavior (like normalization and validation), and handling errors gracefully.

### Key Concepts

- **Structs**: Custom data types to group related fields (e.g., `User`).
- **Struct Tags**: Metadata like `` `json:"name"` `` used by the `encoding/json` package to map JSON fields to struct fields.
- **Methods**: Functions attached to a struct (receiver).
    - **Pointer Receiver (`*User`)**: Used when the method needs to modify the struct instance (e.g., `Normalize`).
    - **Value Receiver (`User`)**: Used when the method doesn't modify the struct or is just reading data (e.g., `Validate`).
- **Data Validation**: Enforcing rules on input data before processing it (checking for empty fields, valid ranges, etc.).
- **Data Normalization**: Cleaning or formatting input data (e.g., trimming whitespace, setting defaults).
- **Error Handling**: Returning and checking `error` values to manage invalid states or failures.

### Example

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

// Normalize sanitizes the user data (Pointer receiver to modify fields)
func (u *User) Normalize() {
    u.Name = strings.TrimSpace(u.Name)
    u.Email = strings.TrimSpace(u.Email)
}

// Validate checks if the user data is valid (Value receiver)
func (u User) Validate() error {
    if u.Age < 18 {
        return errors.New("age must be at least 18")
    }
    return nil
}
```
