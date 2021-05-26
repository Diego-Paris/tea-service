package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	userrepo "github.com/Diego-Paris/tea-service/pkg/models/user_repo"
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

	result, err := userrepo.GetAllUsers()

	if err != nil {
		response := message{"Could not retrieve users. " + err.Error()}
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	//response := message{"Get all users new renewed and amped"}
	respondWithJSON(w, http.StatusOK, result)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	// check if id in path is a valid objectID
	valid := primitive.IsValidObjectID(id)
	if !valid {
		response := message{"User ID is not valid."}
		respondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// search a User by the ID given
	result, err := userrepo.GetUserByID(id)

	if err == mongo.ErrNoDocuments {
		response := message{"User does not exist."}
		respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	if err != nil {
		response := message{"Could not retrieve user. (" + err.Error() + ")"}
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	respondWithJSON(w, http.StatusOK, result)
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
	user, err := userrepo.NewUser(
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

	// check if email is already in use
	emailExists, err := userrepo.IsEmailAlreadyInUse(user.GetEmail())
	if err != nil {
		response := message{"Could not create user. " + err.Error()}
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	if emailExists {
		response := message{"User already exists"}
		respondWithJSON(w, http.StatusConflict, response)
		return
	}

	// save user
	err = userrepo.CreateNew(user)
	if err != nil {
		response := message{"Could not create user. " + err.Error()}
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
