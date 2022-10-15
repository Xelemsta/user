package user

import (
	"user/internal/db"
	"user/internal/transform"
	"user/models"
	"user/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPerPage = int64(10)
	defaultPage    = int64(0)
)

type listUserImpl struct{}

func NewListUserHandler() user.UserListHandler {
	return &listUserImpl{}
}

// Handle implements GET /user.
func (impl *listUserImpl) Handle(params user.UserListParams) middleware.Responder {
	dbUserParams := &db.User{}

	if params.Country != nil {
		dbUserParams.Country = *params.Country
	}

	limit := defaultPerPage
	offset := limit * defaultPage
	dbParams := db.Params{
		Limit:  &limit,
		Offset: &offset,
	}

	if params.PerPage != nil {
		dbParams.Limit = params.PerPage
	}
	if params.Page != nil {
		offset := *params.Page * *params.PerPage
		dbParams.Offset = &offset
	}

	users, err := db.GetUsers(dbUserParams, &dbParams)
	if err != nil {
		log.Errorf(`error while listing users: %+v`, err)

		httpCode := errorToHTTPCode(err)
		return user.NewUserListDefault(httpCode).WithPayload(&models.Error{
			Code:    int64(httpCode),
			Message: err.Error(),
		})
	}

	apiUsers := make([]*models.User, 0, len(users))
	for _, user := range users {
		apiUsers = append(apiUsers, transform.DBUser(user))
	}

	return user.NewUserListOK().WithPayload(&models.GetUserListResponse{
		Users: apiUsers,
		Meta:  nil,
	})
}
