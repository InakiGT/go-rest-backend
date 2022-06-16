package handlers

import (
	"encoding/json"
	"net/http"

	"platzi.com/go/rest-ws/server"
)

type HomeResponse struct {
	Message string `json:"message"` // En go será message pero se serializa a json como message
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // w: Encargado de responderle al cliente, r: Data que envía el cliente
		w.Header().Set("Content-Type", "application/json") // Le dice al cliente que la respuesta es en JSON
		w.WriteHeader(http.StatusOK)                       // Code 200
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Platzi Go",
			Status:  true,
		})
	}
}
