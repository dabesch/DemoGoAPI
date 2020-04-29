package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers - function for registering new controllers
func RegisterControllers() {
	uc := newUserController()

	// Both registrations required
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
