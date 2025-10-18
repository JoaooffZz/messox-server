package handlers

import (
	"encoding/json"
	mb "mb/ports"
	"net/http"
	"strconv"
	wsConn "ws/connection"
)

type Response struct {
	ClientID int `json:"client-id"`
}

type RouteUpgrade struct {
	API_KEY   *string
	HandlerMB mb.HandlerMB
}

func (route *RouteUpgrade) Handler(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "method not is valid", http.StatusMethodNotAllowed)
	// 	return
	// }

	// var response Response
	// err := json.NewDecoder(r.Body).Decode(&response)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{
	// 		"error": "JSON invalid",
	// 	})
	// 	return
	// }
	query := r.URL.Query()
	clientID := query.Get("client_id")
	if clientID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JSON invalid",
		})
	}
	ID, _ := strconv.Atoi(clientID)
	// if err != nil {
	// 	fmt.Println("Erro ao converter:", err)
	// 	return
	// }

	server := wsConn.ServerWS{
		ClientID:  ID,
		HandlerMB: route.HandlerMB,
		W:         w,
		R:         r,
	}
	server.Run()
}
