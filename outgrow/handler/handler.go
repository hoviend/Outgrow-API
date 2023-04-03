package handler

import (
	"outgrow/handler/authhandler"
	"outgrow/handler/userhandler"
)

type RouteHandler struct {
	AuthHandler *authhandler.AuthHandler
	UserHandler *userhandler.UserHandler
}

func NewRouteHandler(
	authHandler *authhandler.AuthHandler,
	userHandler *userhandler.UserHandler,
) *RouteHandler {
	return &RouteHandler{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}
