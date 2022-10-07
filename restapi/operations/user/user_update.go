// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UserUpdateHandlerFunc turns a function with the right signature into a user update handler
type UserUpdateHandlerFunc func(UserUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserUpdateHandlerFunc) Handle(params UserUpdateParams) middleware.Responder {
	return fn(params)
}

// UserUpdateHandler interface for that can handle valid user update params
type UserUpdateHandler interface {
	Handle(UserUpdateParams) middleware.Responder
}

// NewUserUpdate creates a new http.Handler for the user update operation
func NewUserUpdate(ctx *middleware.Context, handler UserUpdateHandler) *UserUpdate {
	return &UserUpdate{Context: ctx, Handler: handler}
}

/*
	UserUpdate swagger:route PUT /v1/user/{userId} user userUpdate

Update an existing user
*/
type UserUpdate struct {
	Context *middleware.Context
	Handler UserUpdateHandler
}

func (o *UserUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUserUpdateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}