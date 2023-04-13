package authhandler

import (
	"errors"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/enum"
	"outgrow/service/authservice"
	"outgrow/service/userservice"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *authservice.AuthService
	UserService *userservice.UserService
}

func NewAuthHandler(authService *authservice.AuthService, userService *userservice.UserService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		UserService: userService,
	}
}

func GetUserInfoFromContext(c *fiber.Ctx) (*dto.UserInfoFromIDToken, error) {
	userInfo := c.Locals(enum.UserInfoKey).(*dto.UserInfoFromIDToken)

	if userInfo == nil {
		return nil, errors.New("user info context get empty value")
	}

	return userInfo, nil
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userInfo, err := GetUserInfoFromContext(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.Validate(userInfo)
	if err != nil {
		errData := dto.Err{
			Message: "error validation",
			Data:    err,
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errData)
	}

	// if data is not exists then insert
	u, err := h.UserService.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			u, err = h.UserService.CreateUser(ctx, userInfo.Email, userInfo.Name)
		}
		if err != nil {
			return err
		}
	}

	resp := dto.SuccessResponse{
		Message: "success",
		Data: dto.AuthLoginResponse{
			UserID: u.ID,
		},
	}

	return c.JSON(resp)
}
