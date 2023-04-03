package handler

import (
	"errors"
	"outgrow/config"
	"outgrow/handler/middleware"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error_message": err.Error(),
			})
		},
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	})

	return app
}

func (h *RouteHandler) SetupRoutes(app *fiber.App, cfg *config.Config) {

	// Middleware
	v1 := app.Group("/v1")

	// Auth
	authAPI := v1.Group("/auth")
	authAPI.Post("/login", middleware.GoogleIDTokenMiddleware(), h.AuthHandler.Login)

	// User
	userAPI := v1.Group("/users", middleware.GoogleIDTokenMiddleware())
	userAPI.Get("/:id", h.UserHandler.GetUser)
	userAPI.Get("/:id/organizations", h.UserHandler.GetUserOrganizations)
	userAPI.Post("/:id/organizations", h.UserHandler.CreateOrganizationByUser)
	userAPI.Post("/:id/organizations/join", h.UserHandler.UserJoinOrganization)

	// handle 404 as fallback
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
