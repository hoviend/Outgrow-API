package eventservice

import (
	"context"
	"errors"
	"outgrow/config"
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/ent/organizationeventtype"
	"outgrow/ent/schema"
	"outgrow/enum"
)

type EventService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewEventService(entClient *ent.Client, config *config.Config) *EventService {
	return &EventService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *EventService) GetOrganizationEventTypeByID(ctx context.Context, id int) (*ent.OrganizationEventType, error) {
	return svc.entClient.OrganizationEventType.Query().
		Where(organizationeventtype.ID(id)).
		Only(ctx)
}

func (svc *EventService) ValidateEventRule(ctx context.Context, eventRules []schema.EventRules) (err error) {
	var (
		totalDebitPercentage,
		totalCreditPercentage float64
	)
	// validate account
	for _, er := range eventRules {
		_, err = svc.entClient.OrganizationAccount.Get(ctx, er.AccountID)
		if err != nil {
			return
		}

		if er.RuleType == enum.EventRulesPercentage {
			if er.TransactionType == enum.TransactionTypeDebit {
				totalDebitPercentage += er.Amount
			} else {
				totalCreditPercentage += er.Amount
			}
		}
	}

	// validate amounts
	switch {
	case totalDebitPercentage < 100:
		return errors.New("invalid total debit percentage")
	case totalCreditPercentage < 100:
		return errors.New("invalid total debit percentage")
	case totalDebitPercentage != totalCreditPercentage:
		return errors.New("mismatch total amount")
	}

	return nil
}

func (svc *EventService) CalculateTotalAmounts(eventRules []schema.EventRules, eventAmount float64) (res dto.TotalEventRuleAmounts) {
	for _, eventRule := range eventRules {
		if eventRule.RuleType == enum.EventRulesPercentage {
			if eventRule.TransactionType == enum.TransactionTypeDebit {
				res.TotalDebitPercentage += eventRule.Amount
			} else {
				res.TotalCreditPercentage += eventRule.Amount
			}
		} else {
			if eventRule.TransactionType == enum.TransactionTypeDebit {
				res.TotalDebitFlatAmount += eventRule.Amount
			} else {
				res.TotalCreditFlatAmount += eventRule.Amount
			}
		}
	}

	// currently, if the flat amount exists then it will reducing the inputted amount first.
	res.TotalDebitAmount = eventAmount - res.TotalDebitFlatAmount
	res.TotalCreditAmount = eventAmount - res.TotalCreditFlatAmount

	return
}
