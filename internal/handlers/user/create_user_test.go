package user_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"user/models"
	"user/testutils"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestCreateUser(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	err := testutils.CleanDB()
	td.CmpNil(t, err)

	t.Run("missing_body", func(t *testing.T) {
		ta.PostJSON(
			"/v1/user",
			nil,
		).CmpStatus(http.StatusBadRequest).
			CmpJSONBody(models.Error{
				Message: "email in body is required",
				Code:    602,
			})
	})

	t.Run("ok", func(t *testing.T) {
		ta.PostJSON(
			"/v1/user",
			json.RawMessage(`{
				"first_name": "abcd",
				"last_name": "efgh",
				"nickname": "azerty",
				"email":"userCreate@test.com",
				"country": "FR",
				"password": "aqwzsxedcrfv"
			}`),
		).CmpStatus(http.StatusCreated).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var user models.User
					errJSON := json.Unmarshal(b, &user)
					if errJSON != nil {
						return errJSON
					}

					td.Cmp(t, user.Country, "FR")
					td.Cmp(t, *user.FirstName, "abcd")
					td.Cmp(t, *user.LastName, "efgh")
					td.Cmp(t, *user.Email, "userCreate@test.com")
					td.Cmp(t, user.Nickname, "azerty")

					return nil
				}),
			)
	})

	t.Run("user_already_exists", func(t *testing.T) {
		ta.PostJSON(
			"/v1/user",
			json.RawMessage(`{
				"first_name": "abcd",
				"last_name": "efgh",
				"nickname": "azerty",
				"email":"userCreate@test.com",
				"country": "FR",
				"password": "aqwzsxedcrfv"
			}`),
		).CmpStatus(http.StatusInternalServerError)
	})
}
