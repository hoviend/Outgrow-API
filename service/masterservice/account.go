package masterservice

import (
	"context"
	"outgrow/ent"
)

func (svc *MasterService) GetMasterAccountTypes(ctx context.Context) ([]*ent.MasterAccountType, error) {
	return svc.entClient.MasterAccountType.Query().
		WithCategories(func(macq *ent.MasterAccountCategoryQuery) { macq.WithAccounts() }).
		All(ctx)
}
