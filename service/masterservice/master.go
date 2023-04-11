package masterservice

import (
	"context"
	"outgrow/config"
	"outgrow/ent"
)

type MasterService struct {
	entClient *ent.Client
	config    *config.Config
}

func NewMasterService(entClient *ent.Client, config *config.Config) *MasterService {
	return &MasterService{
		entClient: entClient,
		config:    config,
	}
}

func (svc *MasterService) GetMasterEventAndAccountData(ctx context.Context) ([]*ent.MasterEventType, []*ent.MasterAccountType, error) {
	masterEventTypes, err := svc.GetMasterEventTypes(ctx)
	if err != nil {
		return nil, nil, err
	}

	masterAccountTypes, err := svc.GetMasterAccountTypes(ctx)
	if err != nil {
		return nil, nil, err
	}

	return masterEventTypes, masterAccountTypes, nil
}
