package handler

import (
	"outgrow/handler/authhandler"
	"outgrow/handler/eventhandler"
	"outgrow/handler/organizationhandler"
	"outgrow/handler/userhandler"
)

type RouteHandler struct {
	AuthHandler         *authhandler.AuthHandler
	UserHandler         *userhandler.UserHandler
	OrganizationHandler *organizationhandler.OrganizationHandler
	EventHandler        *eventhandler.EventHandler
}

func NewRouteHandler(
	authHandler *authhandler.AuthHandler,
	userHandler *userhandler.UserHandler,
	orgHandler *organizationhandler.OrganizationHandler,
	eventHandler *eventhandler.EventHandler,
) *RouteHandler {
	return &RouteHandler{
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		OrganizationHandler: orgHandler,
		EventHandler:        eventHandler,
	}
}
