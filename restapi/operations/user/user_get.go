// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UserGetHandlerFunc turns a function with the right signature into a user get handler
type UserGetHandlerFunc func(UserGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserGetHandlerFunc) Handle(params UserGetParams) middleware.Responder {
	return fn(params)
}

// UserGetHandler interface for that can handle valid user get params
type UserGetHandler interface {
	Handle(UserGetParams) middleware.Responder
}

// NewUserGet creates a new http.Handler for the user get operation
func NewUserGet(ctx *middleware.Context, handler UserGetHandler) *UserGet {
	return &UserGet{Context: ctx, Handler: handler}
}

/*
	UserGet swagger:route GET /v1/user/{userId} user userGet

Get user by it's ID
*/
type UserGet struct {
	Context *middleware.Context
	Handler UserGetHandler
}

func (o *UserGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUserGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
