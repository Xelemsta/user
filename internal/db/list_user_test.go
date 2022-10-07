package db_test

import (
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/td"
)

// GetUsers retrieves users from database
func TestGetUsers(t *testing.T) {
	err := db.Boot()
	td.CmpNil(t, err)
	err = testutils.CleanDB()
	td.CmpNil(t, err)

	firstname1 := "firstname1"
	lastName1 := "lastName1"
	email1 := "getUser1@test.com"
	passwd1 := "passwd1"

	firstname2 := "firstname2"
	lastName2 := "lastName2"
	email2 := "getUser2@test.com"
	passwd2 := "passwd2"

	firstname3 := "firstname3"
	lastName3 := "lastName3"
	email3 := "getUser3@test.com"
	passwd3 := "passwd3"

	dataset := []*db.User{
		{
			FirstName: &firstname1,
			LastName:  &lastName1,
			Email:     &email1,
			Password:  &passwd1,
			Nickname:  "nickname1",
			Country:   "FR",
		},
		{
			FirstName: &firstname2,
			LastName:  &lastName2,
			Email:     &email2,
			Password:  &passwd2,
			Nickname:  "nickname2",
			Country:   "FR",
		},
		{
			FirstName: &firstname3,
			LastName:  &lastName3,
			Email:     &email3,
			Password:  &passwd3,
			Nickname:  "nickname2",
			Country:   "DE",
		},
	}

	for _, d := range dataset {
		err = db.AddUser(d)
		td.CmpNil(t, err)
	}

	t.Run("list_user_ok_default", func(t *testing.T) {
		users, err := db.GetUsers(nil, nil)
		td.CmpNil(t, err)
		td.CmpLen(t, users, 3)
	})

	t.Run("list_user_ok_limit_5", func(t *testing.T) {
		limit := int64(5)
		users, err := db.GetUsers(nil, &db.Params{
			Limit: &limit,
		})
		td.CmpNil(t, err)
		td.CmpLen(t, users, 3)
	})

	t.Run("list_user_ok_limit_5_offet_2", func(t *testing.T) {
		limit := int64(5)
		offset := int64(2)
		users, err := db.GetUsers(nil, &db.Params{
			Offset: &offset,
			Limit:  &limit,
		})
		td.CmpNil(t, err)
		td.CmpLen(t, users, 1)
	})

	t.Run("list_user_ok_limit_1", func(t *testing.T) {
		limit := int64(1)
		users, err := db.GetUsers(nil, &db.Params{
			Limit: &limit,
		})
		td.CmpNil(t, err)
		td.CmpLen(t, users, 1)
	})

	t.Run("list_user_ok_limit_1_offset_5", func(t *testing.T) {
		limit := int64(1)
		offset := int64(5)
		users, err := db.GetUsers(nil, &db.Params{
			Offset: &offset,
			Limit:  &limit,
		})
		td.CmpNil(t, err)
		td.CmpLen(t, users, 0)
	})
}
