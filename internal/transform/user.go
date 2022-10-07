package transform

import (
	"user/internal/db"
	"user/models"

	"github.com/go-openapi/strfmt"
)

// DBUser transforms a db User to its corresponding api model
func DBUser(user *db.User) *models.User {
	apiUser := &models.User{
		Country:   user.Country,
		CreatedAt: strfmt.DateTime(user.CreatedAt),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		UpdatedAt: strfmt.DateTime(user.UpdatedAt),
	}

	if user.ID != nil {
		apiUser.ID = *user.ID
	}

	return apiUser
}
