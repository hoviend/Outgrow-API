package userservice

import (
	"context"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/organization"
	"outgrow/ent/user"

	"github.com/google/uuid"
)

type UserService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewUserService(entClient *ent.Client, config *config.Config) *UserService {
	return &UserService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *UserService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return svc.entClient.User.Query().
		Where(user.Email(email)).
		Only(ctx)
}

func (svc *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return svc.entClient.User.Query().
		Where(user.ID(id)).
		Only(ctx)
}

func (svc *UserService) CreateUser(ctx context.Context, email, name string) (*ent.User, error) {
	return svc.entClient.User.Create().
		SetEmail(email).
		SetName(name).
		Save(ctx)
}

func (svc *UserService) CreateOrganizationByUser(ctx context.Context, name string, user *ent.User) (*ent.Organization, error) {
	return svc.entClient.Organization.Create().
		SetName(name).
		SetPublicID(dto.GeneratePublicID()).
		AddUsers(user).
		Save(ctx)
}

func (svc *UserService) GetUserOrganizations(ctx context.Context, id uuid.UUID, opt dto.GetUserOrganizationsOption) ([]*ent.Organization, dto.PaginateResponse, error) {
	total := 0
	pagination := dto.PaginateResponse{}

	q := svc.entClient.User.Query().
		Where(user.ID(id)).
		QueryOrganizations()

	if opt.OrganizationNameFilter != "" {
		q = q.Where(organization.NameContainsFold(opt.OrganizationNameFilter))
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, dto.PaginateResponse{}, err
	}
	if opt.Paginate != nil {
		q = q.Offset(opt.Paginate.Offset()).
			Limit(opt.Paginate.PerPage)

		pagination = opt.Paginate.ToResponse(total)
	}

	q = q.Order(ent.Desc(organization.FieldCreatedAt))

	orgs, err := q.All(ctx)
	return orgs, pagination, err
}
