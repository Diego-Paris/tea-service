package controllers

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Msg string `json:"msg"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithJPG(w http.ResponseWriter, r *http.Request, imagePath string) {
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, imagePath)
}
