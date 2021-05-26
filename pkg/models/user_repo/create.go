package userrepo

import (
	"github.com/kamva/mgm/v3"
)

// CreateNew inserts a new user into the collection
// Before inserting, checks if there is an already existing
// user with the same email
func CreateNew(newUser *User) error {

	var err error

	// insert into database
	err = mgm.Coll(newUser).Create(newUser)
	if err != nil {
		return err
	}

	return nil
}
