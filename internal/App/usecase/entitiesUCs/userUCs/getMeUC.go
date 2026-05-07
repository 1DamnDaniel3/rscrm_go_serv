package userUCs

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type GetMeUC struct {
	QueryService entitiesrepos.UserAccountQueryService
}

type IGetMeUC interface {
	Execute(ctx context.Context) (*businessobjects.UserBO, error)
}

func NewGetMeUC(QueryService entitiesrepos.UserAccountQueryService) IGetMeUC {
	return &GetMeUC{QueryService}
}

func (uc *GetMeUC) Execute(ctx context.Context) (*businessobjects.UserBO, error) {
	bo, err := uc.QueryService.GetMe(ctx)
	if err != nil {
		return nil, err
	}

	return bo, nil
}
