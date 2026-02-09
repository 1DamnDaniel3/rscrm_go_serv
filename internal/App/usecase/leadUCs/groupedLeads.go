package leadUCs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupedLeadsUC struct {
	repo entitiesrepos.LeadsRepository
}
type IGroupedLeadsUC interface {
	Execute(ctx context.Context, group_id int64, entities *[]entities.Lead) error
}

func NewGroupedLeadsUC(repo entitiesrepos.LeadsRepository) *GroupedLeadsUC {
	return &GroupedLeadsUC{repo}
}

func (uc *GroupedLeadsUC) Execute(ctx context.Context, group_id int64, entities *[]entities.Lead) error {
	if err := uc.repo.GetGroupedLeads(ctx, group_id, entities); err != nil {
		return err
	}
	return nil
}
