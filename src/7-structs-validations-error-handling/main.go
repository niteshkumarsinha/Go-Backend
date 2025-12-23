package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

/*
{
	"name": "John Doe",
	"age": 30,
	"email": "john.doe@example.com"
}
*/

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (u *User) Normalize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	if u.Name == "" {
		u.Name = "John Doe"
	}
	if u.Email == "" {
		u.Email = "john.doe@example.com"
	}
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Age < 18 {
		return errors.New("age must be at least 18")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Normalize()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User created successfully"))
}

func main() {
	fmt.Println("Hello, World!")
	user := User{Name: "Nitesh", Age: 25, Email: "nitesh@example.com"}
	user.Normalize()
	fmt.Println(user)
}
