package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ClientID int `json:"client-id"`
}

type RouteUpgrade struct {
}

func (route *RouteUpgrade) Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not is valid", http.StatusMethodNotAllowed)
		return
	}

	var response Response
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JSON invalid",
		})
		return
	}
}
