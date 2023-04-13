package accountservice

import (
	"context"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/organizationaccount"
	"outgrow/ent/organizationaccountcategory"
	"outgrow/ent/organizationaccounttype"

	"github.com/google/uuid"
)

type AccountService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewAccountService(entClient *ent.Client, config *config.Config) *AccountService {
	return &AccountService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *AccountService) CreateOrganizationAccount(ctx context.Context, payload dto.CreateOrganizationAccountPayload) (*ent.OrganizationAccount, error) {
	return svc.entClient.OrganizationAccount.Create().
		SetCategoryID(payload.AccountCategoryID).
		SetName(payload.AccountName).
		SetCode(payload.Code).
		Save(ctx)
}

func (svc *AccountService) CheckOrganizationAccountCategory(ctx context.Context, organizationID uuid.UUID, catID int) (bool, error) {
	ac, err := svc.entClient.OrganizationAccountCategory.Query().
		Where(organizationaccountcategory.ID(catID)).
		WithType(func(oatq *ent.OrganizationAccountTypeQuery) {
			oatq.Where(organizationaccounttype.OrganizationID(organizationID))
		}).Only(ctx)
	if err != nil {
		return false, err
	}

	return ac != nil, nil
}

func (svc *AccountService) GetAccountCategories(ctx context.Context, opt dto.GetAccountCategoriesOption) ([]*ent.OrganizationAccountCategory, dto.PaginateResponse, error) {
	total := 0
	pagination := dto.PaginateResponse{}

	q := svc.entClient.OrganizationAccountCategory.Query()

	if opt.OrganizationID != uuid.Nil {
		q = q.QueryType().Where(organizationaccounttype.OrganizationID(opt.OrganizationID)).QueryCategories()
	} else {
		q = q.QueryType().QueryCategories()
	}

	if opt.CategoryNameFilter != "" {
		q = q.Where(organizationaccountcategory.NameContainsFold(opt.CategoryNameFilter))
	}

	q = q.WithType()

	total, err := q.Count(ctx)
	if err != nil {
		return nil, dto.PaginateResponse{}, err
	}
	if opt.Paginate != nil {
		q = q.Offset(opt.Paginate.Offset()).
			Limit(opt.Paginate.PerPage)
	}
	pagination = opt.Paginate.ToResponse(total)

	q = q.Order(ent.Asc(organizationaccountcategory.FieldID))

	categories, err := q.All(ctx)
	return categories, pagination, err
}

func (svc *AccountService) GetOrganizationAccountTypes(ctx context.Context, organizationID uuid.UUID, opt dto.GetOrganizationAccountTypesOption) ([]*ent.OrganizationAccountType, dto.PaginateResponse, error) {
	total := 0
	pagination := dto.PaginateResponse{}

	q := svc.entClient.OrganizationAccountType.Query().
		Where(organizationaccounttype.OrganizationID(organizationID))

	if opt.WithCategoriesAndAccounts {
		q = q.WithCategories(func(oacq *ent.OrganizationAccountCategoryQuery) {
			oacq.WithAccounts()
		})
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, dto.PaginateResponse{}, err
	}
	if opt.Paginate != nil {
		q = q.Offset(opt.Paginate.Offset()).
			Limit(opt.Paginate.PerPage)
	}
	pagination = opt.Paginate.ToResponse(total)

	accounts, err := q.All(ctx)
	return accounts, pagination, err
}

func (svc *AccountService) GetOrganizationAccountByID(ctx context.Context, id int) (*ent.OrganizationAccount, error) {
	return svc.entClient.OrganizationAccount.Get(ctx, id)
}

func (svc *AccountService) GetOrganizationAccounts(ctx context.Context, opt dto.GetOrganizationAccountsOption) ([]*ent.OrganizationAccount, dto.PaginateResponse, error) {
	total := 0
	pagination := dto.PaginateResponse{}

	q := svc.entClient.OrganizationAccount.Query().
		WithAccCategory()

	if opt.OrganizationID != uuid.Nil {
		q = q.WithAccCategory(func(oacq *ent.OrganizationAccountCategoryQuery) {
			oacq.WithType(func(oatq *ent.OrganizationAccountTypeQuery) {
				oatq.Where(organizationaccounttype.OrganizationID(opt.OrganizationID))
			})
		})
	} else {
		q = q.WithAccCategory(func(oacq *ent.OrganizationAccountCategoryQuery) {
			oacq.WithType()
		})
	}

	if opt.AccountNameFilter != "" {
		q = q.Where(organizationaccount.NameContainsFold(opt.AccountNameFilter))
	}

	q = q.Order(ent.Asc(organizationaccounttype.FieldID))

	total, err := q.Count(ctx)
	if err != nil {
		return nil, dto.PaginateResponse{}, err
	}
	if opt.Paginate != nil {
		q = q.Offset(opt.Paginate.Offset()).
			Limit(opt.Paginate.PerPage)
	}
	pagination = opt.Paginate.ToResponse(total)

	accounts, err := q.All(ctx)
	return accounts, pagination, err
}
