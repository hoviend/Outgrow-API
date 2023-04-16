package transactionservice

import (
	"context"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/event"
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

	if opt.AccountIDFilter != 0 {
		q = q.Where(
			transaction.AccountID(opt.AccountIDFilter),
		)
	}

	if opt.OrganizationID != uuid.Nil {
		q = q.Where(
			transaction.HasEventWith(event.OrganizationID(opt.OrganizationID)),
		)
	}

	if opt.JoinAccount {
		q = q.WithAccount()
	}

	if opt.JoinEvent {
		q = q.WithEvent(func(eq *ent.EventQuery) { eq.WithType() })
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

	q = q.Order(ent.Desc(transaction.FieldCreatedAt))

	trxs, err := q.All(ctx)
	return trxs, pagination, err
}

func (svc *TransactionService) CreateTransactions(ctx context.Context, transactions []*ent.Transaction) error {

	tcs := make([]*ent.TransactionCreate, 0)
	for _, trx := range transactions {
		t := svc.entClient.Transaction.Create().
			SetAccountID(trx.AccountID).
			SetAmount(trx.Amount).
			SetTransactionDate(trx.TransactionDate).
			SetTransactionType(trx.TransactionType).
			SetEventID(trx.EventID).
			SetNotes(trx.Notes)

		tcs = append(tcs, t)
	}
	return svc.entClient.Transaction.CreateBulk(tcs...).Exec(ctx)
}
