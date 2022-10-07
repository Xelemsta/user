package user

import (
	"context"
	"user/internal/db"
	"user/internal/events"
	"user/internal/transform"
	"user/models"
	"user/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type updateUserImpl struct{}

func NewUpdateUserHandler() user.UserUpdateHandler {
	return &updateUserImpl{}
}

// Handle implements PUT /user
func (impl *updateUserImpl) Handle(params user.UserUpdateParams) middleware.Responder {
	dbUserParams := &db.User{
		ID:        &params.UserID,
		Email:     params.User.Email,
		FirstName: params.User.FirstName,
		LastName:  params.User.LastName,
		Country:   params.User.Country,
		Nickname:  params.User.Nickname,
	}

	updatedUser, err := db.UpdateUser(dbUserParams)
	if err != nil {
		log.Errorf(`error while updating users: %+v`, err)

		httpCode := errorToHTTPCode(err)
		return user.NewUserUpdateDefault(httpCode).WithPayload(&models.Error{
			Code:    int64(httpCode),
			Message: err.Error(),
		})
	}

	log.Infof(`successfully updated user %s`, params.UserID)

	// notify interested services
	go events.Produce(context.Background(), &events.Data{
		UserID: &params.UserID,
		Action: events.UpdateUser,
	})

	return user.NewUserUpdateOK().WithPayload(transform.DBUser(updatedUser))
}
