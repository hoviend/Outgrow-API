package handler

import (
	"outgrow/handler/authhandler"
	"outgrow/handler/organizationhandler"
	"outgrow/handler/userhandler"
)

type RouteHandler struct {
	AuthHandler         *authhandler.AuthHandler
	UserHandler         *userhandler.UserHandler
	OrganizationHandler *organizationhandler.OrganizationHandler
}

func NewRouteHandler(
	authHandler *authhandler.AuthHandler,
	userHandler *userhandler.UserHandler,
	orgHandler *organizationhandler.OrganizationHandler,
) *RouteHandler {
	return &RouteHandler{
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		OrganizationHandler: orgHandler,
	}
}
