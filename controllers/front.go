package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	//in Go, /users and /users/ are treated as if different so we need to
	// explicitly have both
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

/**
creates an encoder and encore the data we receive
*/
func encoreResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
