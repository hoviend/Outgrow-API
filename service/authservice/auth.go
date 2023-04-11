package authservice

import (
	"context"
	"outgrow/config"
	"outgrow/ent"

	"google.golang.org/api/idtoken"
)

type AuthService struct {
	config *config.Config
}

func NewAuthService(entClient *ent.Client, config *config.Config) *AuthService {
	return &AuthService{
		config: config,
	}
}

// DecodeToken.. Parse and validate IDToken from google oauth2 and get the data from it
func (svc *AuthService) DecodeToken(ctx context.Context, token string) (email, name string, err error) {

	payload, err := idtoken.Validate(ctx, token, "")
	if err != nil {
		return "", "", err
	}

	email = payload.Claims["email"].(string)
	name = payload.Claims["name"].(string)

	return email, name, nil
}
