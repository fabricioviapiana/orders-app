package handler

import (
	"encoding/json"
	"net/http"
)

// respondWithJSON centraliza a resposta em JSON e o header de Content-Type
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// respondWithError centraliza o tratamento de erros HTTP
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
