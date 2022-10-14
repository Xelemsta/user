package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // load driver
)

var dbh *sql.DB

// Boot open a connection to postgresql server
// then check if server is alive with a ping.
// Connection is then provided to Dbh which
// will be used by every db functions
func Boot() error {
	var err error

	if err = checkConfig(); err != nil {
		return err
	}

	if dbh == nil {
		connString := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"))

		dbh, err = sql.Open("postgres", connString)
		if err != nil {
			return err
		}

		return dbh.Ping()
	}

	return nil
}

// GetDbh mostly used for test purpose
func GetDbh() *sql.DB {
	return dbh
}

func checkConfig() error {
	if os.Getenv("POSTGRES_HOST") == "" {
		return fmt.Errorf(`POSTGRES_HOST not set`)
	}

	if os.Getenv("POSTGRES_PORT") == "" {
		return fmt.Errorf(`POSTGRES_PORT not set`)
	}

	if os.Getenv("POSTGRES_USER") == "" {
		return fmt.Errorf(`POSTGRES_USER not set`)
	}

	if os.Getenv("POSTGRES_PASSWORD") == "" {
		return fmt.Errorf(`POSTGRES_PASSWORD not set`)
	}

	if os.Getenv("POSTGRES_DB") == "" {
		return fmt.Errorf(`POSTGRES_DB not set`)
	}

	return nil
}
