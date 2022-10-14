package user

import (
	"user/internal/db"
	"user/internal/transform"
	"user/models"
	"user/restapi/operations/user"

	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
)

type getUserImpl struct{}

func NewGetUserHandler() user.UserGetHandler {
	return &getUserImpl{}
}

// Handle implements GET /user/{userId}.
func (impl *getUserImpl) Handle(params user.UserGetParams) middleware.Responder {
	dbUser, err := db.GetUser(params.UserID)
	if err != nil {
		log.Errorf(`error while getting user: %+v`, err)

		httpCode := errorToHTTPCode(err)
		return user.NewUserGetDefault(httpCode).WithPayload(&models.Error{
			Code:    int64(httpCode),
			Message: err.Error(),
		})
	}

	log.Infof(`successfully retrieved user "%s"`, params.UserID)

	return user.NewUserGetOK().WithPayload(transform.DBUser(dbUser))
}
