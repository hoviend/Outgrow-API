package organizationservice

import (
	"context"
	"outgrow/config"
	"outgrow/ent"
	"outgrow/ent/organization"
)

type OrganizationService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewOrganizationService(entClient *ent.Client, config *config.Config) *OrganizationService {
	return &OrganizationService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *OrganizationService) JoinOrganization(ctx context.Context, publicID string, user *ent.User) error {

	org, err := svc.entClient.Organization.Query().
		Where(organization.PublicID(publicID)).
		Only(ctx)
	if err != nil {
		return err
	}

	return org.Update().AddUsers(user).Exec(ctx)
}
