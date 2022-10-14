package testutils

import (
	"user/internal/db"
)

// CleanDB is a convenient function to make sure
// we have an empty db while processing test
func CleanDB() error {
	dbh := db.GetDbh()

	_, err := dbh.Exec(`TRUNCATE TABLE "user";`)
	if err != nil {
		return err
	}

	return nil
}
