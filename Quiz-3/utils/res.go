package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	jadiByte, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Request Failed!", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jadiByte))
}