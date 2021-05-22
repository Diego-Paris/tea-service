package controllers

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func NotFound(w http.ResponseWriter, r *http.Request) {

	response := message{"Route not found."}

	respondWithJSON(w, http.StatusNotFound, response)
}