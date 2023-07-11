package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetMitigations(businessID string) (mitigationOutput []types.Mitigation, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, implemented, created_at FROM risky_public.mitigations(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	mitigationOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Mitigation])
	if err != nil {
		log.Printf("GetMitigations Error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetMitigation(id string) (mitigationOutput types.Mitigation, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, implemented, created_at FROM risky_public.get_mitigation(fn_mitigation_id => $1)", id)
	if err != nil {
		log.Printf("GetMitigation database error: %s", err)
		return
	}

	mitigationOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Mitigation])
	if err != nil {
		log.Printf("GetMitigation parsing error: %s", err)
		return
	}

	return
}

func (m *DBManager) DeleteMitigation(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_mitigation(fn_mitigation_id => $1)", id)
	if err != nil {
		log.Printf("DeleteMitigation error: %s", err)
		return
	}

	return
}

func (m *DBManager) CreateMitigation(mitigationInput types.Mitigation) (mitigationId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_mitigation(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_implemented => $4)`,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.Implemented).Scan(&mitigationId)
	if err != nil {
		log.Printf("CreateMitigation error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateMitigation(mitigationInput types.Mitigation) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_mitigation(
			fn_mitigation_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_implemented => $5)`,
		mitigationInput.ID,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.Implemented)
	if err != nil {
		log.Printf("UpdateMitigation error: %s", err)
		return
	}

	return
}
