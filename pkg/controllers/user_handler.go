package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL)
	fmt.Println("METHOD: ", r.Method)
	log.Println("inside of get all users")
	response := message{"Get all users new renewed and amped"}
	respondWithJSON(w, http.StatusOK, response)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	response := message{"Found user: " + id}
	respondWithJSON(w, http.StatusOK, response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	log.Println("inside of create user")
	response := message{"Posted a new user"}
	respondWithJSON(w, http.StatusOK, response)
}

func GetImage(w http.ResponseWriter, r *http.Request) {

	//TODO add more logic here, like checking and saving edge cases
	path := "/home/paris/workspace/golang/tea-service/images/wack.jpg"
	respondWithJPG(w, r, path)
}
