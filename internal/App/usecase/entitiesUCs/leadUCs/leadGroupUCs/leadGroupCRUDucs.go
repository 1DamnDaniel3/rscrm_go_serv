package leadgroupucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type LeadGroupCRUDucs struct {
	repo entitiesrepos.LeadGroupsRepo
}

type ILeadGroupCRUDucs interface {
	Create(ctx context.Context, leadGroup *entities.LeadGroup) error
	Delete(ctx context.Context, lead_id int64, group_id int64) (*entities.LeadGroup, error)
}

func NewLeadGroupCRUDucs(repo entitiesrepos.LeadGroupsRepo) ILeadGroupCRUDucs {
	return &LeadGroupCRUDucs{repo}
}

// ---=== methods ===---

func (uc *LeadGroupCRUDucs) Create(ctx context.Context, leadGroup *entities.LeadGroup) error {

	if err := uc.repo.Create(ctx, leadGroup); err != nil {
		return err
	}

	return nil
}

func (uc *LeadGroupCRUDucs) Delete(ctx context.Context, lead_id int64, group_id int64) (*entities.LeadGroup, error) {

	relationMap := map[string]any{
		"lead_id":  lead_id,
		"group_id": group_id,
	}
	relation, err := uc.repo.FindRelation(ctx, relationMap)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Delete(ctx, relation.ID, relation); err != nil {
		return nil, err
	}

	return relation, nil
}
