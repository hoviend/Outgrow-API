package transactionservice

import (
	"context"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/organization"
	"outgrow/ent/predicate"
	"outgrow/ent/transaction"

	"github.com/google/uuid"
)

type TransactionService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewTransactionService(entClient *ent.Client, config *config.Config) *TransactionService {
	return &TransactionService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *TransactionService) GetTransactions(ctx context.Context, opt dto.GetTransactionsOption) ([]*ent.Transaction, dto.PaginateResponse, error) {
	total := 0
	pagination := dto.PaginateResponse{}

	q := svc.entClient.Transaction.Query()

	if opt.OrganizationID != uuid.Nil {
		q = q.Where(
			transaction.HasEventWith(predicate.Event(organization.IDEQ(opt.OrganizationID))),
		)
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

	q = q.Order(ent.Desc(transaction.FieldCreatedAt))

	trxs, err := q.All(ctx)
	return trxs, pagination, err
}
