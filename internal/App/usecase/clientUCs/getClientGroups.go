package clientucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetClientGroupUC struct {
	clientQueryService entitiesrepos.ClientsQueryService
}

type IGetClientGroupUC interface {
	Execute(ctx context.Context, client_id int64, groupSlice *[]entities.Group) error
}

func NewGetClientGroupUC(clientQueryService entitiesrepos.ClientsQueryService) *GetClientGroupUC {
	return &GetClientGroupUC{clientQueryService}
}

func (uc *GetClientGroupUC) Execute(ctx context.Context, client_id int64, groupSlice *[]entities.Group) error {
	return uc.clientQueryService.GetClientGroups(ctx, client_id, groupSlice)
}
