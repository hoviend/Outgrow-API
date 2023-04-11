package dto

import "github.com/google/uuid"

type AuthLoginPayload struct {
	IDToken string `json:"id_token" validate:"required"`
}

type AuthLoginResponse struct {
	AccessToken string    `json:"access_token"`
	UserID      uuid.UUID `json:"user_id"`
}

type UserInfoFromIDToken struct {
	Email   string `json:"email" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Subject string `json:"subject"`
}
