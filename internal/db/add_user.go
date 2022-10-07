package db

import (
	"github.com/juju/errors"
)

// AddUser add user in database.
func AddUser(user *User) error {
	if user.FirstName == nil || *user.FirstName == "" {
		return errors.NewBadRequest(nil, "first name is mandatory")
	}

	if user.Email == nil || *user.Email == "" {
		return errors.NewBadRequest(nil, "email is mandatory")
	}

	if user.LastName == nil || *user.LastName == "" {
		return errors.NewBadRequest(nil, "last name is mandatory")
	}

	if user.Password == nil || *user.Password == "" {
		return errors.NewBadRequest(nil, "password is mandatory")
	}

	var userID string

	insertStatement, err := dbh.Prepare(`insert into "user" ("first_name", "last_name", "email", "nickname", "country", "password") values($1, $2, $3, $4, $5, crypt($6, gen_salt('bf'))) RETURNING id`)
	if err != nil {
		return err
	}

	err = insertStatement.QueryRow(
		*user.FirstName,
		*user.LastName,
		*user.Email,
		user.Nickname,
		user.Country,
		user.Password,
	).Scan(&userID)

	if err != nil {
		return err
	}

	// caller will be able to get id from user provided
	user.ID = &userID

	return nil
}
