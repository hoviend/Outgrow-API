package dto

import (
	"outgrow/ent"
	"outgrow/enum"

	"github.com/google/uuid"
)

type CreateOrganizationAccountPayload struct {
	AccountName       string  `json:"account_name" validate:"required"`
	InitialBalance    float64 `json:"initial_balance,omitempty" validate:"omitempty,min=1"`
	AccountCategoryID int     `json:"account_category_id" validate:"required,numeric"`
	Code              string  `json:"code" validate:"required"`
}

type GetAccountCategoryResponse struct {
	DefaultAccountResponse
	AccountType DefaultAccountResponse `json:"account_type"`
}

type GetAccountSelectionResponse struct {
	DefaultAccountResponse
	Code string `json:"code"`
}

type DefaultAccountResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetAccountsParam struct {
	PaginateParam
	Name       string `json:"name,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
}

type GetOrganizationAccountsOption struct {
	Paginate          *PaginateParam
	OrganizationID    uuid.UUID
	AccountNameFilter string
	CategoryID        int
}

type GetOrganizationAccountTypesOption struct {
	Paginate                  *PaginateParam
	WithCategoriesAndAccounts bool
}

type GetAccountCategoriesOption struct {
	Paginate           *PaginateParam
	OrganizationID     uuid.UUID
	CategoryNameFilter string
}

type GetChartofAccountsResponse struct {
	ID              int     `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	CategoryID      int     `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	Balance         float64 `json:"balance"`
	AccountTypeID   int     `json:"account_type_id"`
	AccountTypeName string  `json:"account_type_name"`
}

func (r *GetChartofAccountsResponse) BuildResponse(acc *ent.OrganizationAccount) {
	r.ID = acc.ID
	r.Code = acc.Code
	r.Name = acc.Name
	r.Balance = acc.Balance
	r.CategoryID = acc.CategoryID
	r.CategoryName = acc.Edges.AccCategory.Name
	r.AccountTypeID = acc.Edges.AccCategory.AccountTypeID
	r.AccountTypeName = acc.Edges.AccCategory.Edges.Type.Name
}

type UpdateBalanceParam struct {
	Account         *ent.OrganizationAccount
	Amount          float64
	TransactionType enum.TransactionType
}
