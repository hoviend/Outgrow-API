package main

import (
	"fmt"
	"log"
	"outgrow/config"
	"outgrow/db"
	"outgrow/handler"
	"outgrow/handler/authhandler"
	"outgrow/handler/userhandler"
	authservice "outgrow/service/auth"
	"outgrow/service/organizationservice"
	userservice "outgrow/service/user"
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
	organizationService := organizationservice.NewOrganizationService(client, config)

	// handler
	authHandler := authhandler.NewAuthHandler(authService, userService)
	userHandler := userhandler.NewUserHandler(userService, organizationService)
	routeHandler := handler.NewRouteHandler(authHandler, userHandler)

	routeHandler.SetupRoutes(app, config)

	// run
	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.Port)))
}
