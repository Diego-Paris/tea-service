package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Diego-Paris/tea-service/pkg/models"
)

type UserJSON struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (uj *UserJSON) Decode(requestBody io.ReadCloser) error {
	decoder := json.NewDecoder(requestBody)
	err := decoder.Decode(uj)
	if err != nil {
		return err
	}
	return nil
}

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
	
	var err error

	// Decode request body into struct
	var userJson UserJSON
	err = userJson.Decode(r.Body)
	if err != nil {
		response := message{"Malformed request."}
		respondWithJSON(w, http.StatusBadRequest, response)
		return 
	}


	// Create user obj
	user, err := models.NewUser(
		userJson.FirstName,
		userJson.LastName,
		userJson.Email,
		userJson.Password,
	)
	if err != nil {
		response := message{"Could not create user. " + err.Error()}
		respondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	alreadyExists, err := models.CheckUserExistanceByEmail(user.Email)
	if err != nil {
		response := message{"Could not create user. " + err.Error()}
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	if alreadyExists {
		response := message{"User already exists"}
		respondWithJSON(w, http.StatusConflict, response)
		return
	}

	// save user
	err = user.SaveUser()
	if err != nil {
		response := message{"Could not save user. " + err.Error()}
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := message{"Success"}
	respondWithJSON(w, http.StatusCreated, response)
}

func GetImage(w http.ResponseWriter, r *http.Request) {

	//TODO add more logic here, like checking and saving edge cases
	path := "/home/paris/workspace/golang/tea-service/images/wack.jpg"
	respondWithJPG(w, r, path)
}
