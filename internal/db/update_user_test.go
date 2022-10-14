package db_test

import (
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/td"
)

// UpdateUser updates a user from database
func TestUpdateUser(t *testing.T) {
	err := db.Boot()
	td.CmpNil(t, err)
	err = testutils.CleanDB()
	td.CmpNil(t, err)

	firstname := "firstname"
	lastName := "lastName"
	email := "testUpdateUser@test.com"
	passwd := "passwd"

	user := &db.User{
		Country:   "FR",
		FirstName: &firstname,
		LastName:  &lastName,
		Nickname:  "azerty",
		Email:     &email,
		Password:  &passwd,
	}

	err = db.AddUser(user)
	td.CmpNil(t, err)
	td.CmpNotNil(t, user.ID)

	t.Run("no_params", func(t *testing.T) {
		_, err := db.UpdateUser(nil)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "paramsUser nil")
	})

	t.Run("no_user_id", func(t *testing.T) {
		_, err := db.UpdateUser(&db.User{
			Country:   "FR",
			FirstName: &firstname,
			LastName:  &lastName,
			Nickname:  "azerty",
			Email:     &email,
			Password:  &passwd,
		})
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user id is mandatory")
	})

	t.Run("update_ok", func(t *testing.T) {
		firstname := "firstname"
		lastName := "lastName"
		email := "updateOk@test.com"

		updatedUser, err := db.UpdateUser(&db.User{
			ID:        user.ID,
			Country:   "DE",
			FirstName: &firstname,
			LastName:  &lastName,
			Nickname:  "azertyiop",
			Email:     &email,
		})
		td.CmpNil(t, err)
		td.Cmp(t, *updatedUser.ID, *user.ID)
		td.Cmp(t, updatedUser.Country, "DE")
		td.Cmp(t, *updatedUser.FirstName, firstname)
		td.Cmp(t, *updatedUser.LastName, lastName)
		td.Cmp(t, *updatedUser.Email, email)
	})
}
