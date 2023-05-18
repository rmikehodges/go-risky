package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type MitigationModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	ActionID    uuid.UUID     `json:"actionId" db:"action_id"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetMitigations(businessID string) (mitigationOutput []MitigationModel, err error) {

	rows, err := m.dbPool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented, created_at FROM risky_public.mitigations(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	mitigationOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[MitigationModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetMitigation(id string) (mitigationOutput MitigationModel, err error) {

	rows, err := m.dbPool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented, created_at FROM risky_public.get_mitigation(fn_mitigation_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	mitigationOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[MitigationModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteMitigation(id string) (err error) {

	_, err = m.dbPool.Query(context.Background(), "select risky_public.delete_mitigation(fn_mitigation_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateMitigation(mitigationInput MitigationModel) (err error) {

	_, err = m.dbPool.Query(context.Background(),
		`select risky_public.create_mitigation(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_action_id => $4, 
			fn_implemented => $5)`,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.ActionID,
		mitigationInput.Implemented)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateMitigation(mitigationInput MitigationModel) (err error) {

	_, err = m.dbPool.Query(context.Background(),
		`select risky_public.update_mitigation(
			fn_mitigation_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_action_id => $5, 
			fn_implemented => $6)`,
		mitigationInput.ID,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.ActionID,
		mitigationInput.Implemented)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
