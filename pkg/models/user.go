package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)


// User will store of all our users data
type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"first_name" bson:"first_name"`
	LastName         string `json:"last_name" bson:"last_name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	IsBanned         bool   `json:"is_banned" bson:"is_banned"`
}

// CollectionName returns the name of the collection in the database
func (u *User) CollectionName() string {
	return "my_users"
}

// NewUser creates a new user
func NewUser(firstName, lastName, email, password string) (*User, error) {


	var err error
	newUser := &User{}

	err = newUser.SetFirstName(firstName)
	if err != nil {
		return nil, err
	}

	err = newUser.SetLastName(lastName)
	if err != nil {
		return nil, err
	}

	err = newUser.SetEmail(email)
	if err != nil {
		return nil, err
	}

	err = newUser.SetPassword(password)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) SaveUser() error {

	err := mgm.Coll(u).Create(u)

	if err != nil {
		return err
	}

	return nil
}

/*

---- All property setting methods ----

*/

func (u *User) SetFirstName(firstName string) error {

	minLength := 1
	maxLength := 32

	// check minimum length
	if len(firstName) < minLength {
		msg := fmt.Sprintf("First name cannot be less than %v characters.", minLength)
		return errors.New(msg)
	}

	// check max length
	if len(firstName) > maxLength {
		msg := fmt.Sprintf("First name cannot be more than %v characters.", maxLength)
		return errors.New(msg)
	}

	u.FirstName = firstName

	return nil
}

func (u *User) SetLastName(lastName string) error {

	minLength := 1
	maxLength := 64

	// check minimum length
	if len(lastName) < minLength {
		msg := fmt.Sprintf("Last name cannot be less than %v characters.", minLength)
		return errors.New(msg)
	}

	// check max length
	if len(lastName) > maxLength {
		msg := fmt.Sprintf("Last name cannot be more than %v characters.", maxLength)
		return errors.New(msg)
	}

	u.LastName = lastName

	return nil
}

func (u *User) SetEmail(email string) error {

	minLength := 1
	maxLength := 64

	// check minimum length
	if len(email) < minLength {
		msg := fmt.Sprintf("Email cannot be less than %v characters.", minLength)
		return errors.New(msg)
	}

	// check max length
	if len(email) > maxLength {
		msg := fmt.Sprintf("Email cannot be more than %v characters.", maxLength)
		return errors.New(msg)
	}

	u.Email = email

	return nil
}

func (u *User) SetPassword(password string) error {

	minLength := 1
	maxLength := 32

	// check minimum length
	if len(password) < minLength {
		msg := fmt.Sprintf("Password cannot be less than %v characters.", minLength)
		return errors.New(msg)
	}

	// check max length
	if len(password) > maxLength {
		msg := fmt.Sprintf("Password cannot be more than %v characters.", maxLength)
		return errors.New(msg)
	}

	u.Password = password

	return nil
}

/*

---- All find and filter methods ----

*/


func CheckUserExistanceByEmail(email string) (bool, error) {
	// Get the document's collection
	user := &User{}
	coll := mgm.Coll(user)

	_ = coll.First(bson.M{"email":"chosenone23@gmail.com"}, user)

	count, err := coll.CountDocuments(mgm.Ctx(), bson.M{})
	log.Printf("%v -- %v", count, err)


	log.Println("HERE WE ARE")
	log.Println(user)

	return false, nil
}
