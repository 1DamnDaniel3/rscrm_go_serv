package studentclientUCs

import (
	"context"

	clientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_policies"
	studentclientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_client_policies"
	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateRelationUC struct {
	tx         services.Transaction
	scRepo     entitiesrepos.StudentClientsRepo
	clientRepo entitiesrepos.ClientRepo

	studentClientPolicies studentclientpolicies.IStudentClientCrudPolicy
	clientPolicies        clientpolicies.IClientCrudPolicy
}

type ICreateRelationUC interface {
	Execute(ctx context.Context, studentClient entities.StudentClient, bo *businessobjects.GetStudentClientsBO) error
}

func NewCreateRelationUC(
	tx services.Transaction,
	scRepo entitiesrepos.StudentClientsRepo,
	clientRepo entitiesrepos.ClientRepo,
	studentClientPolicies studentclientpolicies.IStudentClientCrudPolicy,
	clientPolicies clientpolicies.IClientCrudPolicy,
) *CreateRelationUC {
	return &CreateRelationUC{tx, scRepo, clientRepo, studentClientPolicies, clientPolicies}
}

func (uc *CreateRelationUC) Execute(ctx context.Context,
	studentClient entities.StudentClient, bo *businessobjects.GetStudentClientsBO) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {

		scCreateScope, err := uc.studentClientPolicies.CanCreate(txCtx)
		if err != nil {
			return err
		}

		if err := uc.scRepo.Create(txCtx, &studentClient, scCreateScope); err != nil {
			return err
		}

		clientCreateScope, err := uc.clientPolicies.CanCreate(txCtx)
		if err != nil {
			return err
		}

		client := entities.Client{}
		if err := uc.clientRepo.GetByID(txCtx, studentClient.Client_id, &client, clientCreateScope); err != nil {
			return err
		}

		bo.Relation_id = studentClient.ID
		bo.Client = client
		bo.Is_payer = studentClient.Is_payer
		bo.Relation = studentClient.Relation

		return nil
	})
}
