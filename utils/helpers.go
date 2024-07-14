package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// get a path param from request
func GetPathParam(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	value, ok := vars[name]
	if !ok {
		return "", fmt.Errorf("invalid or missing %s in request param", name)
	}
	return value, nil
}
