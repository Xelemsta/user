package user

import (
	"context"
	"time"
	"user/internal/db"
	"user/internal/events"
	"user/internal/transform"
	"user/models"
	"user/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type createUserImpl struct{}

func NewCreateUserHandler() user.UserCreateHandler {
	return &createUserImpl{}
}

// Handle implements POST /user.
func (impl *createUserImpl) Handle(params user.UserCreateParams) middleware.Responder {
	dbUser := &db.User{
		Country:   params.User.Country,
		CreatedAt: time.Time(params.User.CreatedAt),
		Email:     params.User.Email,
		FirstName: params.User.FirstName,
		LastName:  params.User.LastName,
		Nickname:  params.User.Nickname,
		UpdatedAt: time.Time(params.User.UpdatedAt),
		Password:  params.User.Password,
	}

	err := db.AddUser(dbUser)
	if err != nil {
		log.Errorf(`error while adding user: %+v`, err)

		httpCode := errorToHTTPCode(err)
		return user.NewUserCreateDefault(httpCode).WithPayload(&models.Error{
			Code:    int64(httpCode),
			Message: err.Error(),
		})
	}

	log.Infof(`successfully created user with uuid: "%s"`, *dbUser.ID)

	// notify interested services
	go events.Produce(context.Background(), &events.Data{
		UserID: dbUser.ID,
		Action: events.CreateUserAction,
	})

	return user.NewUserCreateCreated().WithPayload(transform.DBUser(dbUser))
}
