package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type message struct {
	Msg string `json:"msg"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	response := message{"Get all users"}

	respondWithJSON(w, http.StatusOK, response)

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	response := message{"Found user: " + id}
	respondWithJSON(w, http.StatusConflict, response)
}
