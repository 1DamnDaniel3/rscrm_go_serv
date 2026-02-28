package studentclientsucs

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateRelationUC struct {
	tx         services.Transaction
	scRepo     entitiesrepos.StudentClientsRepo
	clientRepo entitiesrepos.ClientRepo
}

type ICreateRelationUC interface {
	Execute(ctx context.Context, studentClient entities.StudentClient, bo *businessobjects.GetStudentClientsBO) error
}

func NewCreateRelationUC(
	tx services.Transaction,
	scRepo entitiesrepos.StudentClientsRepo,
	clientRepo entitiesrepos.ClientRepo) *CreateRelationUC {
	return &CreateRelationUC{tx, scRepo, clientRepo}
}

func (uc *CreateRelationUC) Execute(ctx context.Context,
	studentClient entities.StudentClient, bo *businessobjects.GetStudentClientsBO) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {

		if err := uc.scRepo.Create(txCtx, &studentClient); err != nil {
			return err
		}

		client := entities.Client{}
		if err := uc.clientRepo.GetByID(txCtx, studentClient.Client_id, &client); err != nil {
			return err
		}

		bo.Relation_id = studentClient.ID
		bo.Client = client
		bo.Is_payer = studentClient.Is_payer
		bo.Relation = studentClient.Relation

		return nil
	})
}
