package clientucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupedClientsUC struct {
	clientRepo entitiesrepos.ClientRepo
}

type IGroupedClientsUC interface {
	Execute(ctx context.Context, group_id int64, entities *[]entities.Client) error
}

func NewGroupedClientsUC(
	clientRepo entitiesrepos.ClientRepo) *GroupedClientsUC {
	return &GroupedClientsUC{clientRepo}
}

func (uc *GroupedClientsUC) Execute(ctx context.Context, group_id int64, entities *[]entities.Client) error {
	return uc.clientRepo.GetGroupedClients(ctx, group_id, entities)
}
