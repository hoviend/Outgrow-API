package masterservice

import (
	"context"
	"outgrow/ent"
)

func (svc *MasterService) GetMasterEventTypes(ctx context.Context) ([]*ent.MasterEventType, error) {
	return svc.entClient.MasterEventType.Query().
		All(ctx)
}
