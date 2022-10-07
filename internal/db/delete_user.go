package db

import (
	"fmt"

	"github.com/juju/errors"
)

// RemoveUser deletes a user from database.
func RemoveUser(userID string) error {
	if userID == "" {
		return errors.NewBadRequest(nil, "user id is mandatory")
	}

	// is user exists ?
	_, err := GetUser(userID)
	if err != nil {
		return err
	}

	result, err := dbh.Exec(`DELETE FROM "user" WHERE id=$1`, userID)
	if err != nil {
		return err
	}

	nb, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if nb == 0 {
		return fmt.Errorf(`no row affected while trying to delete user %s`, userID)
	}

	return err
}
