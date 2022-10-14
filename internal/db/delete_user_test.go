package db_test

import (
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/td"
)

func TestRemoveUser(t *testing.T) {
	err := db.Boot()
	td.CmpNil(t, err)
	err = testutils.CleanDB()
	td.CmpNil(t, err)

	t.Run("no_user_id", func(t *testing.T) {
		err := db.RemoveUser("")
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user id is mandatory")
	})

	t.Run("user_not_found", func(t *testing.T) {
		err := db.RemoveUser("34b4691c-4db9-4269-a053-1cd082ccaff7")
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user not found")
	})

	t.Run("delete_ok", func(t *testing.T) {
		firstname := "abcd"
		lastName := "efgh"
		email := "deleteUser@test.com"
		passwd := "aqwzsxedcrfv"

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

		err = db.RemoveUser(*user.ID)
		td.CmpNil(t, err)

		err := db.RemoveUser(*user.ID)
		td.CmpNotNil(t, err)
		td.Cmp(t, err.Error(), "user not found")
	})
}
