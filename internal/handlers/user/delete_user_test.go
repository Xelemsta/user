package user_test

import (
	"fmt"
	"net/http"
	"testing"
	"user/internal/db"
	"user/testutils"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestDeleteUser(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	err := testutils.CleanDB()
	td.CmpNil(t, err)

	firstname := "abcd"
	lastName := "efgh"
	email := "a.b@c.d"
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

	t.Run("delete_id_not_found", func(t *testing.T) {
		ta.DeleteJSON(
			"/v1/user/34b4691c-4db9-4269-a053-1cd082ccaff7",
			nil,
		).CmpStatus(http.StatusNotFound)
	})

	t.Run("delete_ok", func(t *testing.T) {
		ta.DeleteJSON(
			fmt.Sprintf("/v1/user/%s", *user.ID),
			nil,
		).CmpStatus(http.StatusNoContent).OrDumpResponse()
	})

	t.Run("delete_id_not_found_2", func(t *testing.T) {
		ta.DeleteJSON(
			fmt.Sprintf("/v1/user/%s", *user.ID),
			nil,
		).CmpStatus(http.StatusNotFound)
	})
}
