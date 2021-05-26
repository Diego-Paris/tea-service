package userrepo

import (
	"errors"
	"fmt"

	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"first_name" bson:"first_name"`
	LastName         string `json:"last_name" bson:"last_name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	IsBanned         bool   `json:"is_banned" bson:"is_banned"`
}

type Users []User

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

	newUser.SetBannedStatus(false)

	return newUser, nil
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

func (u *User) SetBannedStatus(status bool) {
	u.IsBanned = status
}

/*

---- All property getting methods ----

*/

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetBannedStatus() bool {
	return u.IsBanned
}

// String is the stringer method ------ //
func (u *User) String() string {
	return fmt.Sprintf(
		"{\n_id: %v,\nfirst_name: %v,\nlast_name: %v,\nemail: %v,\npassword: %v,\nis_banned: %v,\n}",
		u.ID,
		u.GetFirstName(),
		u.GetLastName(),
		u.GetEmail(),
		u.GetPassword(),
		u.GetBannedStatus(),
	)
}
