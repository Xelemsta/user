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

func TestUpdateUser(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	err := testutils.CleanDB()
	td.CmpNil(t, err)

	firstname := "firstname"
	lastName := "lastName"
	email := "TestUpdateUser@test.fr"
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

	t.Run("update_no_body", func(t *testing.T) {
		ta.PutJSON(
			fmt.Sprintf("/v1/user/%s", *user.ID),
			nil,
		).CmpStatus(http.StatusBadRequest)
	})

	t.Run("update_user_not_found", func(t *testing.T) {
		ta.PutJSON(
			"/v1/user/34b4691c-4db9-4269-a053-1cd082ccaff7",
			json.RawMessage(`{
				"first_name": "firstname1",
				"last_name": "lastName1",
				"nickname": "azerty",
				"email":"e.f@g.h",
				"country": "DE"
			}`),
		).CmpStatus(http.StatusNotFound)
	})

	var userUpdated models.User

	t.Run("update_user_ok", func(t *testing.T) {
		ta.PutJSON(
			fmt.Sprintf("/v1/user/%s", *user.ID),
			json.RawMessage(`{
				"first_name": "firstname1",
				"last_name": "lastName1",
				"nickname": "azerty",
				"email":"updateUser@test.com",
				"country": "DE"
			}`),
		).CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					errJSON := json.Unmarshal(b, &userUpdated)
					if errJSON != nil {
						return errJSON
					}

					td.Cmp(t, *userUpdated.FirstName, "firstname1")
					td.Cmp(t, *userUpdated.LastName, "lastName1")
					td.Cmp(t, *userUpdated.Email, "updateUser@test.com")
					td.Cmp(t, userUpdated.Country, "DE")
					return nil
				}),
			)
	})

	t.Run("get_user_updated", func(t *testing.T) {
		ta.Get(
			fmt.Sprintf("/v1/user/%s", userUpdated.ID),
		).CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var getUser models.User
					errJSON := json.Unmarshal(b, &getUser)
					if errJSON != nil {
						return errJSON
					}

					td.Cmp(t, *getUser.FirstName, *userUpdated.FirstName)
					td.Cmp(t, *getUser.LastName, *userUpdated.LastName)
					td.Cmp(t, *getUser.Email, *userUpdated.Email)
					td.Cmp(t, getUser.Country, userUpdated.Country)
					return nil
				}),
			)
	})
}
