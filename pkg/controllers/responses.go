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

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {

	response := message{"Method not allowed."}

	respondWithJSON(w, http.StatusMethodNotAllowed, response)
}

// TODO : Add proper error logging when these critical mistakes occurr 
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				response := message{"An error ocurred in the server."}
				respondWithJSON(w, http.StatusInternalServerError, response)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
