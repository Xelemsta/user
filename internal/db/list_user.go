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

	// XXX: count rows for pagination purposes

	return buildAndExecQuery(userParams, limit, offset)
}

func buildAndExecQuery(userParams *User, limit int64, offset int64) ([]*User, error) {
	var whereClause string
	var values []interface{}
	if userParams != nil {
		if userParams.Country != "" {
			whereClause = ` WHERE "country" = $1`
			values = append(values, userParams.Country)
		}
	}

	values = append(values, limit, offset)

	rootQuery := `SELECT
		"country", "created_at", "email", "first_name", "id", "last_name", "nickname", "updated_at"
		FROM "user"`

	if whereClause != "" {
		rootQuery += whereClause + ` ORDER BY "id" DESC LIMIT $2 OFFSET $3`
	} else {
		rootQuery += ` ORDER BY "id" DESC LIMIT $1 OFFSET $2`
	}

	stmt, err := dbh.Prepare(rootQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(values...)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
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

	return users, nil
}
