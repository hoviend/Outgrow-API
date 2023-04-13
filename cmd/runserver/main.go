package main

import (
	"fmt"
	"log"
	"outgrow/config"
	"outgrow/db"
	"outgrow/handler"
	"outgrow/handler/authhandler"
	"outgrow/handler/eventhandler"
	"outgrow/handler/organizationhandler"
	"outgrow/handler/userhandler"
	"outgrow/service/accountservice"
	"outgrow/service/authservice"
	"outgrow/service/eventservice"
	"outgrow/service/masterservice"
	"outgrow/service/organizationservice"
	"outgrow/service/transactionservice"
	"outgrow/service/userservice"
)

func main() {
	// TODO: dep inject library
	config := config.GetConfig()

	app := handler.NewFiberApp()

	client := db.NewDBClient(*config)
	defer client.Close()

	// TODO: check env if not local doesnt have to debug
	client = client.Debug()

	// svc
	authService := authservice.NewAuthService(client, config)
	userService := userservice.NewUserService(client, config)
	masterService := masterservice.NewMasterService(client, config)
	organizationService := organizationservice.NewOrganizationService(client, config)
	accountService := accountservice.NewAccountService(client, config)
	transactionService := transactionservice.NewTransactionService(client, config)
	eventService := eventservice.NewEventService(client, config)

	// handler
	authHandler := authhandler.NewAuthHandler(authService, userService)
	userHandler := userhandler.NewUserHandler(userService, organizationService, masterService)
	organizationHandler := organizationhandler.NewOrganizationHandler(organizationService, accountService, transactionService, eventService)
	eventHandler := eventhandler.NewEventHandler(eventService, accountService)

	routeHandler := handler.NewRouteHandler(
		authHandler,
		userHandler,
		organizationHandler,
		eventHandler,
	)

	routeHandler.SetupRoutes(app, config)

	// run
	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.Port)))
}
