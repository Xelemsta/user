package db_test

import (
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/td"
)

func TestAddUser(t *testing.T) {
	err := db.Boot()
	td.CmpNil(t, err)
	err = testutils.CleanDB()
	td.CmpNil(t, err)

	firstname := "firstname"
	lastname := "lastname"
	email := "addUser@test.com"
	password := "passwd"

	t.Run("no_first_name", func(t *testing.T) {
		user := db.User{
			LastName: &lastname,
			Country:  "FR",
			Email:    &email,
			Password: &password,
			Nickname: "nickname",
		}

		err := db.AddUser(&user)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "first name is mandatory")
		td.CmpNil(t, user.ID)
	})

	t.Run("no_last_name", func(t *testing.T) {
		user := db.User{
			FirstName: &firstname,
			Country:   "FR",
			Email:     &email,
			Password:  &password,
			Nickname:  "nickname",
		}

		err := db.AddUser(&user)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "last name is mandatory")
		td.CmpNil(t, user.ID)
	})

	t.Run("no_email", func(t *testing.T) {
		user := db.User{
			LastName:  &lastname,
			FirstName: &firstname,
			Country:   "FR",
			Password:  &password,
			Nickname:  "nickname",
		}

		err := db.AddUser(&user)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "email is mandatory")
		td.CmpNil(t, user.ID)
	})

	t.Run("no_password", func(t *testing.T) {
		user := db.User{
			LastName:  &lastname,
			FirstName: &firstname,
			Country:   "FR",
			Email:     &email,
			Nickname:  "nickname",
		}

		err := db.AddUser(&user)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "password is mandatory")
		td.CmpNil(t, user.ID)
	})

	t.Run("ok", func(t *testing.T) {
		user := db.User{
			FirstName: &firstname,
			LastName:  &lastname,
			Country:   "FR",
			Email:     &email,
			Password:  &password,
			Nickname:  "nickname",
		}

		err := db.AddUser(&user)
		td.CmpNil(t, err)
		td.CmpNotNil(t, user.ID)
		td.CmpNotEmpty(t, *user.ID)
	})
}
