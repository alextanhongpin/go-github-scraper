package encoder

import (
	"encoding/json"
	"net/http"
)

// JSON returns the encoded response as json
func JSON(w http.ResponseWriter, err error, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
