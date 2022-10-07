package user_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"user/internal/db"
	"user/models"
	"user/testutils"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestGetUser(t *testing.T) {
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

	t.Run("get_user_not_found", func(t *testing.T) {
		ta.Get("/v1/user/34b4691c-4db9-4269-a053-1cd082ccaff7").CmpStatus(http.StatusNotFound)
	})

	t.Run("get_user_ok", func(t *testing.T) {
		ta.Get(fmt.Sprintf("/v1/user/%s", *user.ID)).CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var user models.User
					errJSON := json.Unmarshal(b, &user)
					if errJSON != nil {
						return errJSON
					}

					td.Cmp(t, user.Country, "FR")
					td.Cmp(t, *user.FirstName, firstname)
					td.Cmp(t, *user.LastName, lastName)
					td.Cmp(t, *user.Email, email)
					td.Cmp(t, user.Nickname, "azerty")

					return nil
				}),
			)
	})

	err = db.RemoveUser(*user.ID)
	td.CmpNil(t, err)

	t.Run("get_user_not_found_after_delete", func(t *testing.T) {
		ta.Get(fmt.Sprintf("/v1/user/%s", *user.ID)).CmpStatus(http.StatusNotFound)
	})

}
