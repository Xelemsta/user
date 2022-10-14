package db_test

import (
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/td"
)

// GetUser retrieves user from database
func TestGetUser(t *testing.T) {
	err := db.Boot()
	td.CmpNil(t, err)
	err = testutils.CleanDB()
	td.CmpNil(t, err)

	t.Run("no_user_id", func(t *testing.T) {
		_, err := db.GetUser("")
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user id is mandatory")
	})

	t.Run("user_not_found", func(t *testing.T) {
		_, err := db.GetUser("34b4691c-4db9-4269-a053-1cd082ccaff7")
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user not found")
	})

	t.Run("get_user_ok", func(t *testing.T) {
		firstname := "abcd"
		lastName := "efgh"
		email := "getUser@test.com"
		passwd := "aqwzsxedcrfv"

		createdUser := &db.User{
			Country:   "FR",
			FirstName: &firstname,
			LastName:  &lastName,
			Nickname:  "azerty",
			Email:     &email,
			Password:  &passwd,
		}

		err = db.AddUser(createdUser)
		td.CmpNil(t, err)
		td.CmpNotNil(t, createdUser.ID)

		getUser, err := db.GetUser(*createdUser.ID)
		td.CmpNil(t, err)

		td.Cmp(t, getUser.Country, "FR")
		td.Cmp(t, *getUser.FirstName, firstname)
		td.Cmp(t, *getUser.LastName, lastName)
		td.Cmp(t, getUser.Nickname, "azerty")
		td.Cmp(t, *getUser.Email, email)
	})
}
