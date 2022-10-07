package transform_test

import (
	"testing"
	"time"
	"user/internal/db"
	"user/internal/transform"

	"github.com/maxatome/go-testdeep/td"
)

func TestTransformUser(t *testing.T) {
	now := time.Now()
	email := "email"
	firstname := "firstname"
	lastname := "lastname"
	nickname := "nickname"
	id := "34b4691c-4db9-4269-a053-1cd082ccaff7"

	dbUser := &db.User{
		ID:        &id,
		Country:   "FR",
		CreatedAt: now,
		Email:     &email,
		FirstName: &firstname,
		LastName:  &lastname,
		Nickname:  nickname,
		UpdatedAt: now,
	}

	modelsUser := transform.DBUser(dbUser)
	td.CmpNotNil(t, modelsUser)
	td.Cmp(t, modelsUser.ID, *dbUser.ID)
	td.Cmp(t, modelsUser.Country, dbUser.Country)
	td.Cmp(t, time.Time(modelsUser.CreatedAt), dbUser.CreatedAt)
	td.Cmp(t, *modelsUser.Email, *dbUser.Email)
	td.Cmp(t, *modelsUser.FirstName, *dbUser.FirstName)
	td.Cmp(t, *modelsUser.LastName, *dbUser.LastName)
	td.Cmp(t, modelsUser.Nickname, dbUser.Nickname)
	td.Cmp(t, time.Time(modelsUser.UpdatedAt), dbUser.UpdatedAt)
}
