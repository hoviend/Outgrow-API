package organizationservice

import (
	"context"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/organization"
	"outgrow/ent/organizationeventtype"

	"github.com/google/uuid"
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

func (svc *OrganizationService) CopyMasters(
	ctx context.Context,
	organizationID uuid.UUID,
	masterEventTypes []*ent.MasterEventType,
	masterAccountTypes []*ent.MasterAccountType,
) error {
	//TODO: would be better if run in 1 transaction?
	orgEventTypes := make([]*ent.OrganizationEventTypeCreate, 0)
	orgAccounts := make([]*ent.OrganizationAccountCreate, 0)

	for _, at := range masterAccountTypes {
		oat, err := svc.entClient.OrganizationAccountType.Create().
			SetOrganizationID(organizationID).
			SetName(at.Name).
			Save(ctx)

		if err != nil {
			return err
		}

		for _, ac := range at.Edges.Categories {
			oac, err := svc.entClient.OrganizationAccountCategory.Create().
				SetTypeID(oat.ID).
				SetName(ac.Name).
				SetDescription(ac.Description).Save(ctx)

			if err != nil {
				return err
			}

			for _, a := range ac.Edges.Accounts {
				oa := svc.entClient.OrganizationAccount.Create().
					SetCategoryID(oac.ID).
					SetName(a.Name)

				orgAccounts = append(orgAccounts, oa)
			}

			_, err = svc.entClient.OrganizationAccount.CreateBulk(orgAccounts...).
				Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	for _, et := range masterEventTypes {
		oet := svc.entClient.OrganizationEventType.Create().
			SetName(et.Name).
			SetDescription(et.Description).
			SetRules(et.Rules).
			SetOrganizationID(organizationID)

		orgEventTypes = append(orgEventTypes, oet)
	}

	_, err := svc.entClient.OrganizationEventType.CreateBulk(orgEventTypes...).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (svc *OrganizationService) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*ent.Organization, error) {
	return svc.entClient.Organization.Query().
		Where(organization.ID(id)).
		Only(ctx)
}

func (svc *OrganizationService) GetOrganizationEventTypes(ctx context.Context, id uuid.UUID) ([]*ent.OrganizationEventType, error) {
	return svc.entClient.OrganizationEventType.Query().
		Where(organizationeventtype.OrganizationID(id)).
		All(ctx)
}

func (svc *OrganizationService) GetOrganizationEventTypeByID(ctx context.Context, id int, organizationID uuid.UUID) (*ent.OrganizationEventType, error) {
	return svc.entClient.OrganizationEventType.Query().
		Where(organizationeventtype.ID(id)).
		Where(organizationeventtype.OrganizationID(organizationID)).
		Only(ctx)
}

func (svc *OrganizationService) GetOrganizationByParam(ctx context.Context, opt dto.GetOrganizationOption) (*ent.Organization, error) {
	q := svc.entClient.Organization.Query()

	if opt.ID != uuid.Nil {
		q = q.Where(organization.ID(opt.ID))
	}

	if opt.WithEvents {
		q = q.WithEvents()
	}

	return q.Only(ctx)
}

func (svc *OrganizationService) CreateEvent(ctx context.Context, payload dto.CreateOrganizationEventPayload) (*ent.Event, error) {
	return svc.entClient.Event.Create().
		SetAmount(payload.Amount).
		SetEventTypeID(payload.EventTypeID).
		SetOrganizationID(payload.OrganizationID).
		Save(ctx)
}

func (svc *OrganizationService) CreateOrganizationEventType(ctx context.Context, payload dto.CreateOrganizationEventTypePayload) (*ent.OrganizationEventType, error) {
	return svc.entClient.OrganizationEventType.Create().
		SetName(payload.Name).
		SetDescription(payload.Description).
		SetOrganizationID(payload.OrganizationID).
		SetRules(payload.Rules).
		Save(ctx)
}
