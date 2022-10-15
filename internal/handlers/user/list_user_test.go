package user_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"user/internal/db"
	"user/models"
	"user/testutils"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestListUser(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	err := testutils.CleanDB()
	td.CmpNil(t, err)

	firstname1 := "firstname1"
	lastName1 := "lastName1"
	email1 := "a.b@c.d"
	passwd1 := "passwd1"

	firstname2 := "firstname2"
	lastName2 := "lastName2"
	email2 := "c.d@x.y"
	passwd2 := "passwd2"

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
			Country:   "DE",
		},
	}

	for _, d := range dataset {
		err = db.AddUser(d)
		td.CmpNil(t, err)
	}

	t.Run("list_user_default", func(t *testing.T) {
		ta.Get("/v1/user").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 2)
					return nil
				}),
			).OrDumpResponse()
	})

	t.Run("list_user_per_page_5", func(t *testing.T) {
		ta.Get("/v1/user?per_page=5").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 2)
					return nil
				}),
			).OrDumpResponse()
	})

	t.Run("list_user_per_page_5_page_2", func(t *testing.T) {
		ta.Get("/v1/user?per_page=5&page=2").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 0)
					return nil
				}),
			)
	})

	t.Run("list_user_per_page_1", func(t *testing.T) {
		ta.Get("/v1/user?per_page=1").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 1)
					return nil
				}),
			)
	})

	t.Run("list_user_per_page_5_country_FR", func(t *testing.T) {
		ta.Get("/v1/user?per_page=5&country=FR").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 1)
					return nil
				}),
			)
	})

	t.Run("list_user_per_page_5_country_DE", func(t *testing.T) {
		ta.Get("/v1/user?per_page=5&country=DE").CmpStatus(http.StatusOK).
			CmpJSONBody(
				td.Code(func(b json.RawMessage) error {
					var userListResponse models.GetUserListResponse
					errJSON := json.Unmarshal(b, &userListResponse)
					if errJSON != nil {
						return errJSON
					}

					td.CmpLen(t, userListResponse.Users, 1)
					return nil
				}),
			)
	})
}
