package encoder

import (
	"encoding/json"
	"net/http"
)

// ErrorJSON represents the error payload in json format
type ErrorJSON struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// SuccessJSON represents the success payload in json format
type SuccessJSON struct {
	Data interface{} `json:"data,omitempty"`
}

// JSON returns the encoded response as json
func JSON(w http.ResponseWriter, err error, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		errJSON := ErrorJSON{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(errJSON)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(SuccessJSON{v}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
