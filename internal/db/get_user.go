package db

import "github.com/juju/errors"

// GetUser retrieves a user from database.
func GetUser(userID string) (*User, error) {
	if userID == "" {
		return nil, errors.NewBadRequest(nil, "user id is mandatory")
	}

	selectStatement, err := dbh.Prepare(
		`SELECT "country", "created_at", "email", "first_name",
		"id", "last_name", "nickname", "updated_at" FROM "user" where id = $1`,
	)

	if err != nil {
		return nil, err
	}

	user := User{}
	err = selectStatement.QueryRow(userID).Scan(
		&user.Country,
		&user.CreatedAt,
		&user.Email,
		&user.FirstName,
		&user.ID,
		&user.LastName,
		&user.Nickname,
		&user.UpdatedAt,
	)

	if err != nil {
		// catch up user not found to let caller handle it properly
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.NewNotFound(nil, "user not found")
		}
	}

	return &user, err
}
