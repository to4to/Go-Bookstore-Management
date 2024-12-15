package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the body of an HTTP request and unmarshals it into the provided interface.
// If an error occurs during reading or unmarshalling, the function simply returns without any action.
//
// Parameters:
//   - r: The HTTP request from which to read the body.
//   - x: The interface into which the body should be unmarshalled.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
