package user

import (
	"context"
	"user/internal/db"
	"user/internal/events"
	"user/models"
	"user/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type deleteUserImpl struct{}

func NewDeleteUserHandler() user.UserDeleteHandler {
	return &deleteUserImpl{}
}

// Handle implements DELETE /user.
func (impl *deleteUserImpl) Handle(params user.UserDeleteParams) middleware.Responder {
	err := db.RemoveUser(params.UserID)
	if err != nil {
		log.Errorf(`error while deleting user: %+v`, err)

		httpCode := errorToHTTPCode(err)
		return user.NewUserDeleteDefault(httpCode).WithPayload(&models.Error{
			Code:    int64(httpCode),
			Message: err.Error(),
		})
	}

	log.Infof(`successfully deleted user "%s"`, params.UserID)

	// notify interested services
	go events.Produce(context.Background(), &events.Data{
		UserID: &params.UserID,
		Action: events.DeleteUser,
	})

	return user.NewUserDeleteNoContent()
}
