package util

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON is a helper to return the response as JSON
func ResponseJSON(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
