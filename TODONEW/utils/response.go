package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseOK(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if body == nil {
		w.Write([]byte("Successfully Done !"))
	} else {
		json.NewEncoder(w).Encode(body)
	}

}

func ResponseError(w http.ResponseWriter, error *APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.StatusCode)
	jsonValue, _ := json.Marshal(error)
	w.Write(jsonValue)
}
