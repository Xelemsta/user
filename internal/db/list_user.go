package db

const (
	limitDefault  = int64(10)
	offsetDefault = int64(0)
)

// GetUsers retrieves a list of user from database.
func GetUsers(userParams *User, params *Params) ([]*User, error) {
	limit := limitDefault
	offset := offsetDefault
	if params != nil {
		if params.Limit != nil {
			limit = *params.Limit
		}

		if params.Offset != nil {
			offset = *params.Offset
		}
	}

	// XXX: add filter from userParams
	// XXX: count rows for pagination purposes

	selectStatement, err := dbh.Prepare(
		`SELECT "country", "created_at", "email", "first_name", "id", "last_name", "nickname", "updated_at"
			FROM "user"
			ORDER BY "id" DESC
			OFFSET $1 LIMIT $2`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := selectStatement.Query(offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(
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
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
