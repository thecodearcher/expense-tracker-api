package auth

import "fmt"

import "net/http"

// Login authenticates user into application
func Login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Test")
}
