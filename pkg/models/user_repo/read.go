package userrepo

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers() (Users, error) {

	var result Users

	err := mgm.Coll(&User{}).SimpleFind(&result, bson.M{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserByID(userID string) (*User, error) {
	// Get the document's collection
	user := &User{}
	coll := mgm.Coll(user)

	// Find and decode the doc to a book model.
	err := coll.FindByID(userID, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func getCountByFilter(filter interface{}) (int, error) {

	user := &User{}
	coll := mgm.Coll(user)

	totalFound, err := coll.CountDocuments(mgm.Ctx(), filter)

	if err != nil {
		return -1, err
	}

	return int(totalFound), nil
}

func IsEmailAlreadyInUse(email string) (bool, error) {

	filter := bson.M{"email": email}

	total, err := getCountByFilter(filter)

	if err != nil {
		return false, err
	}

	return total > 0, nil
}
