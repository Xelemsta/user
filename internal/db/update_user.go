package db

import (
	"github.com/juju/errors"
)

// UpdateUser updates a user from database.
func UpdateUser(paramsUser *User) (*User, error) {
	if paramsUser == nil {
		return nil, errors.NewBadRequest(nil, `paramsUser nil`)
	}
	if paramsUser.ID == nil {
		return nil, errors.NewBadRequest(nil, `user id is mandatory`)
	}

	user, err := GetUser(*paramsUser.ID)
	if err != nil {
		return nil, err
	}

	user.Country = paramsUser.Country
	user.Nickname = paramsUser.Nickname

	if paramsUser.FirstName != nil {
		user.FirstName = paramsUser.FirstName
	}

	if paramsUser.LastName != nil {
		user.LastName = paramsUser.LastName
	}

	if paramsUser.Email != nil {
		user.Email = paramsUser.Email
	}

	updateStatement, err := dbh.Prepare(`UPDATE "user" SET 
		first_name = $1,
		last_name = $2,
		nickname = $3,
		country = $4,
		email = $5
		WHERE id = $6`,
	)
	if err != nil {
		return nil, err
	}

	_, err = updateStatement.Exec(*user.FirstName, *user.LastName, user.Nickname, user.Country, *user.Email, *user.ID)
	if err != nil {
		return nil, err
	}

	// reload user to make sure we have fresh data
	updatedUser, err := GetUser(*paramsUser.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
