package middleware

import (
	"outgrow/dto"
	"outgrow/enum"
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/idtoken"
)

func GoogleIDTokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error_message": "Authorization header is missing",
			})
		}
		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if bearerToken == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error_message": "Invalid Authorization header format",
			})
		}
		payload, err := idtoken.Validate(c.UserContext(), bearerToken, "")
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error_message": "Failed to decode auth token",
			})
		}

		userInfo := &dto.UserInfoFromIDToken{
			Email:   payload.Claims["email"].(string),
			Name:    payload.Claims["name"].(string),
			Subject: payload.Subject,
		}
		c.Locals(enum.UserInfoKey, userInfo)

		return c.Next()
	}
}
